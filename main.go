package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	rnm "github.com/pitakill/rickandmortyapigowrapper"
)

var templateText = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<style>
	*{
		margin: 0;
		padding: 0;
		box-sizing: border-box;
	}
	
	body{
		width: 100vw;
		height: 100vh;
		display: flex;
		gap: 2rem;
		flex-wrap: wrap;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 20px;
		font-family: Georgia, 'Times New Roman', Times, serif;
		background-color: #333;
		
	}
	h1{
		position: relative;
		color: transparent;
	}
	h1:after{
		position: absolute;
		top: 0;
		left: 40px;
		width: 100%;
		height: 100%;
		content: 'Rick and Morty';
		color: #0070ff;
		animation: animateTitle 5s linear infinite;
	}
	h1 span{
		color: #fff;
	}
	#container{
		display: flex;
		flex-direction: column;
		justify-content: flex-start;
		align-items: flex-start;
		background-color: #fff;
		width: 300px;
		height: auto;
		border-radius: 10px;
	}
	.characters{
		position: relative;
		width: 300px;
		height: 300px;
		opacity: 1;
		transition: 0.3s;
		border-top-left-radius: 10px;
		border-top-right-radius: 10px;
		display: flex;
		justify-content: flex-end;
		align-items: center;
		cursor: pointer;
		background-size: contain;
		background-repeat: no-repeat;
		background-image: url({{.Image}});
	}
	.characters:hover{
		transition: 0.3s;
	}
	.title, .type, .status{
		display: flex;
		margin: 0px 10px;
		flex-direction: column;
		color: #333;
		font-size: 2rem;
		text-transform: uppercase;
		letter-spacing: 0.5px;
		font-weight: 600;
	}
	.title span, .type span, .status span{
		color: rgba(0,0,0,0.5);
		text-transform: capitalize;
		font-size: 1.8rem;
		font-weight: 500;
	}
	@keyframes animateTitle{
		0%{
			filter: hue-rotate(0deg);
		}
		100%{
			filter: hue-rotate(360deg);
		}
	}
	</style>
    <title>Rick And Morty</title>
</head>
<body>
	<h1><span>&#128125;</span> Rick and Morty <span>&#128299;</span></h1>
    <div id="container">
        <div class="characters"></div>
        <div class="title">
            name:
            <span>{{.Name}}</span>
        </div>
        <div class="type">
            type:
            <span>{{.Type}}</span>
        </div>
        <div class="status">
            status:
            <span>{{.Status}}</span>
        </div>
    </div>
    
</body>
</html>

`

func getCharacter(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("character").Parse(templateText)
	if err != nil {
		log.Fatal(err)
	}
	vars := mux.Vars(r)
	charId, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
		log.Println(charId)
	}
	character, err := rnm.GetCharacter(charId)
	if err != nil {
		log.Fatal(err)
	}
	c := struct {
		Image  string
		Name   string
		Type   string
		Status string
	}{
		Image:  character.Image,
		Name:   character.Name,
		Type:   character.Species,
		Status: character.Status,
	}
	// fmt.Fprint(w, character)
	t.ExecuteTemplate(w, "character", c)

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/character/{id}", getCharacter)
	http.ListenAndServe(":8000", router)
}
