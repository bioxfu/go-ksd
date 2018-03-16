package main

import (
	"database/sql"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB

// Fields must start with a upper case
type PsiteResource struct {
	PsiteID	string `json:"psite_id"`
	ProteinID	string `json:"protein_id"`
	GeneID	string `json:"gene_id"`
	Position	int `json:"position"`
	Amino_acid	string `json:"amino_acid"`
	Probability	float64 `json:"probability"`
	Class	string `json:"class"`
	Window	string `json:"window"`
}

func getPsite(c *gin.Context) {
	var x PsiteResource
	id := c.Param("protein_id")
	rows, err := DB.Query("select psiteID, proteinID, geneID, position, amino_acid, probability, class, window from psites where proteinID=?", id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		var X[] PsiteResource
		for rows.Next() {
			err := rows.Scan(&x.PsiteID, &x.ProteinID, &x.GeneID, &x.Position, &x.Amino_acid, &x.Probability, &x.Class, &x.Window)
			if err != nil {
				log.Fatal(err)
			}
			X = append(X, x)
		}
		log.Println(X)
		c.JSON(200, gin.H{
			"results": X,
		})
	}
}

//

type AliasResource struct {
	GeneID	string `json:"gene_id"`
	Aliases	string `json:"aliases"`
}

func getAlias(c *gin.Context) {
	var x AliasResource
	id := c.Param("gene_id")
	rows, err := DB.Query("select geneID, aliase from aliases where geneID=?", id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		var X[] AliasResource
		for rows.Next() {
			err := rows.Scan(&x.GeneID, &x.Aliases)
			if err != nil {
				log.Fatal(err)
			}
			X = append(X, x)
		}
		log.Println(X)
		c.JSON(200, gin.H{
			"results": X,
		})
	}
}

//

type DomainResource struct {
	ProteinID	string `json:"protein_id"`
	Source	string `json:"source"`
	Domain	string `json:"domain"`
	Start	int `json:"start"`
	End	int `json:"end"`
}

func getDomain(c *gin.Context) {
	var x DomainResource
	id := c.Param("protein_id")
	rows, err := DB.Query("select proteinID, source, domain, start, end from domains where proteinID=?", id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		var X[] DomainResource
		for rows.Next() {
			err := rows.Scan(&x.ProteinID, &x.Source, &x.Domain, &x.Start, &x.End)
			if err != nil {
				log.Fatal(err)
			}
			X = append(X, x)
		}
		log.Println(X)
		c.JSON(200, gin.H{
			"results": X,
		})
	}
}

//

type DescriptionResource struct {
	ProteinID	string `json:"protein_id"`
	Biotype	string `json:"biotype"`
	Description	string `json:"description"`
}

func getDescription(c *gin.Context) {
	var x DescriptionResource
	id := c.Param("protein_id")
	rows, err := DB.Query("select proteinID, biotype, description from descriptions where proteinID=?", id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		var X[] DescriptionResource
		for rows.Next() {
			err := rows.Scan(&x.ProteinID, &x.Biotype, &x.Description)
			if err != nil {
				log.Fatal(err)
			}
			X = append(X, x)
		}
		log.Println(X)
		c.JSON(200, gin.H{
			"results": X,
		})
	}
}

//

type GoResource struct {
	GeneID	string `json:"gene_id"`
	GoTerm	string `json:"GO_term"`
	GoAccession	string `json:"GO_accession"`
	GoDomain	string `json:"GO_domain"`
}

func getGO(c *gin.Context) {
	var x GoResource
	id := c.Param("gene_id")
	rows, err := DB.Query("select geneID, GO_term, GO_accession, GO_domain from GO where geneID=?", id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		var X[] GoResource
		for rows.Next() {
			err := rows.Scan(&x.GeneID, &x.GoTerm, &x.GoAccession, &x.GoDomain)
			if err != nil {
				log.Fatal(err)
			}
			X = append(X, x)
		}
		log.Println(X)
		c.JSON(200, gin.H{
			"results": X,
		})
	}
}

//

func main() {
	var err error
	DB, err = sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	r := gin.Default()
	// add routes to REST verbs
	r.GET("/v1/psite/:protein_id", getPsite)
	r.GET("/v1/alias/:gene_id", getAlias)
	r.GET("/v1/domain/:protein_id", getDomain)
	r.GET("/v1/description/:protein_id", getDescription)
	r.GET("/v1/GO/:gene_id", getGO)
	r.Run(":8080")

}
