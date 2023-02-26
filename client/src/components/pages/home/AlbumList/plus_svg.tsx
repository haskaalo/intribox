import Buttonize from "@home/propsbuilder/buttonize";
import { CreateNewEmptyAlbum } from "@home/request/album";
import * as React from "react";

// This is how it looks: https://jsfiddle.net/gd639jvb/
function PlusSVG() {
  async function handleClick() {
    const value = window.prompt("Album name") // TODO: Change this
    
    // Create empty album
    try {
      const id = await CreateNewEmptyAlbum({name: value});
      alert(`Successfully created new album ID: ${id}`) // TODO: Change this to a better signifiant
    } catch (err) {
      alert("Error while creating new album") // TODO: Change this to a better signifiant
      console.error(err);
    }
  }
  
  const {tabIndex, onClick, onKeyDown} = Buttonize(handleClick);

   return <div style={{cursor: "pointer"}} role="button" tabIndex={tabIndex} onClick={onClick} onKeyDown={onKeyDown}>
    <svg width="100%" height="100%" viewBox="0 0 500 500" xmlns="http://www.w3.org/2000/svg" xmlnsXlink="http://www.w3.org/1999/xlink">
        <line x1="100" y1="250" x2="400" y2="250" stroke="white" strokeWidth="20" />
        <line x1="250" y1="100" x2="250" y2="400" stroke="white" strokeWidth="20" />
    </svg>
  </div> 
}

export default PlusSVG;