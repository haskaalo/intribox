import { addMedia, Media} from "@home/redux/slice/mediagrid";
import * as React from "react";
import { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { GetListMedia } from "@home/request/media";
import { AppDispatch, RootState } from "@home/redux/store";
import { Row, Col } from "reactstrap";
import MediaComponent from "./media_component";
import AlbumList from "../AlbumList";

function MediaGrid() {
    const medias = useSelector((state: RootState) => state.mediagrid.loadedMedias);
    const dispatch: AppDispatch = useDispatch();

    async function loadMedia() {
        try {
            // Fetch first 25 medias metadata, of course we'll change that
            const response = await GetListMedia();
            const addedResponse: Media[] = response.map((media: Media) => {
                media.downloaded = false;
                return media;
            });

            // Add the list of media metadata to the media grid so it can start
            // preloading
            dispatch(addMedia(addedResponse));
        } catch(err) {
            // TODO: Change that
            alert("Error while loading medias check console")
            // eslint-disable-next-line no-console
            console.error(err);
        }
    }

    // componentDidMount
    useEffect(() => {
        loadMedia();
    }, []);

    const mediaRows= [];

    for (let i = 0; i < medias.length; i++) {
        const mediaComponent = <Col xs={3} md={2} className="media-col"><MediaComponent media={medias[i]} /></Col>;
        if (i % 12 === 0) {
            mediaRows.push([mediaComponent]);
        } else {
            mediaRows[mediaRows.length - 1].push(mediaComponent);
        }
    }

    let keyIndex = 0
    return <div>
        <AlbumList />
        <div className="section-media-grid">
        {mediaRows.map(rowChildrens => <Row key={`row-${keyIndex++}`}>{rowChildrens}</Row>)}
        </div>
    </div>
}

export default MediaGrid
