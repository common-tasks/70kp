import React from "react";

class AboutEvents extends React.Component{

    HandleButtonClick = (e)=>{
        e.preventDefault();
        console.log('I am clicked' +e);
    }

render() {
    return (
         <div>
            <h1>Handling events</h1>
            <button onClick={this.HandleButtonClick} >click</button>
         </div>
    );
}
}

export default AboutEvents;