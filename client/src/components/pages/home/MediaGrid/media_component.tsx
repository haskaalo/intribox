import { Media, mediaBeenDownloaded } from "@home/redux/slice/mediagrid";
import React from "react";
import { connect } from "react-redux";

export interface IProps {
    media: Media;
    setMediaDownloaded: typeof mediaBeenDownloaded;
}

function MediaComponent(props: IProps) {
    const { media, setMediaDownloaded } = props;

    // TODO: add placeholder image?
    return <div className="media-wrap">
        <img src={media.download_url} alt="" loading="lazy" onLoad={() => setMediaDownloaded(media.id)}/>
    </div>
}

export default connect(null, {setMediaDownloaded: mediaBeenDownloaded})(MediaComponent);