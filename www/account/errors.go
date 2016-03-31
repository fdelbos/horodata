package account

import (
	"dev.hyperboloide.com/fred/horodata/html"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetError(c *gin.Context, err error) {
	log.WithFields(log.Fields{
		"error": err,
	}).Error("WWW account error detected.")

	c.Writer.WriteHeader(http.StatusInternalServerError)

	data := map[string]interface{}{
		"Title":  "Erreure du serveur",
		"Error":  "Une erreure a été détectée!",
		"Detail": "Veuillez recommencer."}
	html.Render("account/error.html", c, data, http.StatusInternalServerError)
}

func UserNotVerified(c *gin.Context) {
	data := map[string]interface{}{
		"Title":  "Connexion refusée",
		"Error":  "L'adresse email associée a ce compte n' a pas été verifiée.",
		"Detail": "Verifiez cette adresse aupres du fournisseur de votre compte et recommencez."}
	html.Render("account/error.html", c, data, http.StatusInternalServerError)
}
