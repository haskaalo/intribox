import * as React from "react";
import AlbumBlock from "./album";
import PlusSVG from "./plus_svg";

function AlbumList() {
    return <div className="albums-row">
        <AlbumBlock>
            <PlusSVG/>
        </AlbumBlock>
    </div>
}

export default AlbumList;