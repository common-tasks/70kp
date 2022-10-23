import React from "react";

class RequestAnalyzer extends React.Component {
    state = {
        reqCount :this.props.counter,
    }
    render() {
        return (
             <>
             <h1>{this.state.reqCount}</h1>
             </>
        );
    }

}

export default RequestAnalyzer