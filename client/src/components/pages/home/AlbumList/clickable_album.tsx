import * as React from "react";
import AlbumBlock from "./album_block";

export interface AlbumBlockProps {
    albumID: string;
}

function ClickableAlbum(props: AlbumBlockProps) {

    const { albumID } = props;
    console.log(albumID);
    
    return <AlbumBlock>
        <div style={{width: "100%", height: "100%", flex: 1, backgroundColor: "gray"}} />
    </AlbumBlock>
}

export default ClickableAlbum;
