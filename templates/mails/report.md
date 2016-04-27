Rapport du {{ .date }}
---

Donnée          | Total
---             | ---
Utilisateurs    | {{ .total_users}}
Groupes         | {{ .total_groups }}
Invités         | {{ .total_guests }}
Saisies         | {{ .total_jobs }}

### Plans:
10 Utilisateurs | 30 Utilisateurs | 100 Utilisateurs
---             | ---             | ---
{{ .total_plans.small }} | {{ .total_plans.medium }} | {{ .total_plans.large }}

### Nouveaux utilisateurs:
Email   | Nom   | id
---     |---    |---
{{ range .new_users }}[{{ .email }}](mailto:{{ .email }}) | {{ .name }} | {{ .id }}
{{ end }}
