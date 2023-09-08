import React from "react";
import "./TextArea.css";

export function TextArea(props) {
    return(
            <textarea className="text-area" id= {props.id} name={props.name} cols={props.cols} rows={props.rows}>{props.value}</textarea>
    )    
}