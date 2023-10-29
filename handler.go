package main

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"strings"

	xmltoyaml "github.com/Prasang-money/go-parser/xmlToYaml"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// Handler for /ping endpoint
func health(c *gin.Context) {
	c.String(http.StatusOK, "server is running")
}

// Handler for /xml/yaml endpoint
func xmlToYaml(c *gin.Context) {

	// Reading request body in in memoey buffer
	var buf bytes.Buffer
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Creating a new decoder from buf(xml) data. which we will use to traverse the given xml data.
	dec := xml.NewDecoder(strings.NewReader(buf.String()))
	mp, err := xmltoyaml.XmlToMap("", []xml.Attr{}, dec)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Marshaling the map of type map[string]interface{} into yaml data
	yamlData, err := yaml.Marshal(mp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// writing the response
	c.Header("Content-Type", "application/x-yaml")
	c.String(http.StatusOK, string(yamlData))

}

// Handler for /xml/isValid endpoint
func isXml(c *gin.Context) {
	var buf bytes.Buffer

	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = xml.Unmarshal(buf.Bytes(), new(interface{}))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}
