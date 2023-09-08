import React from 'react'
import './Input.css'

export function Input(props) {
    return (<input className={props.class} type={props.type} id={props.id} onChange={props.onChange} />)
}