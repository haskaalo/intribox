import * as React from "react";

function AlbumBlock(props: {children: React.ReactNode}) {
    const {children} = props;

    return <div className="album-block">
        {children}
    </div>
}

export default AlbumBlock;