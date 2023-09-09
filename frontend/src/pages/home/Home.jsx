import React, { useState, useEffect } from "react";
import axios from "axios";
import "./Home.css";
import { Button } from "../../components/button/Button";
import { TextArea } from "../../components/textArea/TextArea";
import { Input } from "../../components/Input/Input";

export function Home() {
    const [archivo, setArchivo] = useState("");
    const [consola, setConsola] = useState("");
    const api = axios.create({
        baseURL: `http://localhost:4200/`
    })

    useEffect(() => {
        document.getElementsByName("consola")[0].value = consola;
      }, [consola]);

    useEffect(() => {
        document.getElementsByName("texto1")[0].value = archivo;
    }, [archivo])   

    function addConsola(newConsola) {
        setConsola(newConsola);
        console.log(consola);
    }

    async function peticion() {
        const cadena = document.getElementsByName("texto1")[0].value
        await sendGetRequest(cadena);
    }

    function sendGetRequest(cadena) {
        new Promise(()=> {
            api.get('/gramatica', {
                params: {
                    gramatica: cadena
                }
            }).then(res => {
                // res.data
                addConsola(res.data)
            }).catch(err => console.log(err))
        })
    }
    // aqui va la lectura para el archivo
    async function leerArchivo(e) {
        const file = e.target.files[0];
        const reader = new FileReader();
        reader.readAsText(file);
        reader.onload = () => {
            setArchivo(reader.result)
        }
        reader.onerror = () => {
            console.log('file error', reader.error)
        }
    }


    return(<div className="Home-div">
        <div className="bloque-botones">
            <div className="file-system">
                <label for="upload-photo" class="boton-normal">Archivo</label>
                <Input class="Input-normal" type="file" id="upload-photo" onChange={leerArchivo} />
            </div>
        </div>

        <div className="bloque">
            <TextArea name="texto1" className="text-area" cols="80" rows="28"/>
            <textarea name="consola" className="text-area" cols="80" rows="28">{consola}</textarea>
        </div>
        <div className="bloque">
            <Button onClick={peticion} value="compile"/>
        </div>
    </div>)
}