# -*- coding: utf-8 -*-

import os, csv, StringIO, sys, xlsxwriter
from flask import Flask, request, send_file
from dateutil.parser import parse

app = Flask(__name__)

if 'PRODUCTION_MODE' in os.environ:
    app.debug = False
else:
    app.debug = True

@app.route('/', methods=['POST'])
def gen_xlsx():

    js = request.json
    si = StringIO.StringIO()
    workbook = xlsxwriter.Workbook(si)
    dateFormat = workbook.add_format({'num_format': 'dd/mm/yyyy'})
    currencyFormat = workbook.add_format({'num_format': '0.00'})
    hourFormat = workbook.add_format({'num_format': '0.00'})

    guests_time = {}
    guests_cost = {}

    tasks_time = {}
    tasks_cost = {}

    customers_time = {}
    customers_cost = {}

    def do_sheet(name, colName, times, costs):
        sheet = workbook.add_worksheet(name)
        sheet.set_column('A:A', 25)
        sheet.set_column('B:B', 12)
        sheet.set_column('C:C', 12)
        sheet.set_column('D:D', 12)

        sheet.write("A1", colName)
        sheet.write("B1", u"Durée Totale (en heures)")
        sheet.write("C1", u"Durée Totale (en minutes)")
        sheet.write("D1", u"Coût Total")

        line = 1
        for i, duration in times.items():
            line += 1
            sheet.write("A%s" % line, i)
            sheet.write("B%s" % line, duration / 3600.0, hourFormat)
            sheet.write("C%s" % line, duration / 60)
            sheet.write("D%s" % line, costs[i], currencyFormat)

        sheet.write("A%s" % (line + 2), u"Total")
        sheet.write_formula("B%s" % (line + 2),"=SUM(B2:B%s)" % line)
        sheet.write_formula("C%s" % (line + 2),"=SUM(C2:C%s)" % line)
        sheet.write_formula("D%s" % (line + 2),"=SUM(D2:D%s)" % line)

        chart = workbook.add_chart({'type': 'pie'})
        chart.add_series({
            'name': 'Repartition du temps (en heures)',
            'categories': '=%s!$A$2:$A$%s' % (name, line),
            'values': '=%s!$B$2:$B$%s' % (name, line),
        })
        chart.set_title({'name': "Repartition du temps"})
        sheet.insert_chart("A%s" % (line + 4), chart)

    try:

        saisies = workbook.add_worksheet("Saisies")
        saisies.set_column('A:A', 12)
        saisies.set_column('B:B', 25)
        saisies.set_column('C:C', 25)
        saisies.set_column('D:D', 25)
        saisies.set_column('E:E', 12)
        saisies.set_column('F:F', 12)
        saisies.set_column('G:G', 12)
        saisies.set_column('H:H', 50)

        saisies.write("A1", u"Date")
        saisies.write("B1", u"Utilisateur")
        saisies.write("C1", u"Dossier")
        saisies.write("D1", u"Tâche")
        saisies.write("E1", u"Durée (en heures)")
        saisies.write("F1", u"Durée (en minutes)")
        saisies.write("G1", u"Coût")
        saisies.write("H1", u"Commentaire")

        line = 1
        for row in js:
            line += 1
            creation = parse(row["created"]).replace(tzinfo=None)
            guest = row["creator"]
            task = row["task"]
            customer = row["customer"]

            saisies.write_datetime("A%s" % line, creation, dateFormat)
            saisies.write("B%s" % line, guest)
            saisies.write("C%s" % line, customer)
            saisies.write("D%s" % line, task)
            saisies.write("E%s" % line, row["duration"] / 3600.0, hourFormat)
            saisies.write("F%s" % line, row["duration"] / 60)
            saisies.write("G%s" % line, row["cost"], currencyFormat)
            saisies.write("H%s" % line, row["comment"])

            if guest not in guests_time:
                guests_time[guest] = 0
                guests_cost[guest] = 0.0
            guests_time[guest] += row["duration"]
            guests_cost[guest] += row["cost"]

            if task not in tasks_time:
                tasks_time[task] = 0
                tasks_cost[task] = 0.0
            tasks_time[task] += row["duration"]
            tasks_cost[task] += row["cost"]

            if customer not in customers_time:
                customers_time[customer] = 0
                customers_cost[customer] = 0.0
            customers_time[customer] += row["duration"]
            customers_cost[customer] += row["cost"]

        do_sheet(u"Utilisateurs", u"Utilisateur", guests_time, guests_cost)
        do_sheet(u"Dossiers", u"Dossier", customers_time, customers_cost)
        do_sheet(u"Tâches", u"Tâche", tasks_time, tasks_cost)

    finally:
        workbook.close()

    si.seek(0)
    return send_file(si)


if __name__ == '__main__':
    app.run(host='0.0.0.0')
