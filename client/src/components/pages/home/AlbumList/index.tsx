import { addAlbums } from "@home/redux/slice/albumlist";
import { AppDispatch, RootState } from "@home/redux/store";
import { KnownError } from "@home/request";
import { GetListAlbum } from "@home/request/album";
import * as React from "react";
import { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import AlbumBlock from "./album_block";
import ClickableAlbum from "./clickable_album";
import PlusSVG from "./plus_svg";

function AlbumList() {
    const albums = useSelector((state: RootState) => state.albumlist.albumlist);
    const dispatch: AppDispatch = useDispatch();
    
    async function loadAlbums() {
        try {
            const response = await GetListAlbum();
            
            dispatch(addAlbums(response));
        } catch(err) {
            if (err === KnownError.UNAUTHORIZED) return;

            alert("Error while loading medias check console") // TODO: Change that
            // eslint-disable-next-line no-console
            console.error(err);
        }
    }

    // Acts like componentDidMount
    useEffect(() => {
        loadAlbums();
    }, []);

    // Prepare rending

    return <div className="albums-row">
        <AlbumBlock>
            <PlusSVG/>
        </AlbumBlock>
        {albums.map((album) => <ClickableAlbum albumID={album.id} key={`album-${album.id}`}/>)}
    </div>
}

export default AlbumList;