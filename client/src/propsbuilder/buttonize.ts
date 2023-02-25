// This function allows you to "buttonize" a div or whatever that isn't a button or link
// the correct way.
// NOTE: Require role "button" props to the element.
function Buttonize(onClickHandler: React.MouseEventHandler) {
    return {
        tabIndex: 0,
        onClick: onClickHandler,
        onKeyDown: (event: any) => {
            if (event.key === "Enter") onClickHandler(event);
        }
    }
}

export default Buttonize;
