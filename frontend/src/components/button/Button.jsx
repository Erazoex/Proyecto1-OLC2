import React from "react";
import "./Button.css";

export function Button(props) {
    return(
    <button className="boton-normal" type={props.type}  onClick={props.onClick}>{props.value}</button>
    )
}