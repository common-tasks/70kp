import React from "react";

class ConditionalRendereing extends React.Component {
    constructor(props){
        super(props);
        this.isLoggedIn = props.isLoggedIn;
    }
render() {
    if(this.isLoggedIn)
    {
        return (<this.User />)
    }
    else
    {
        return <this.GuestUser />
    }
    
}

User = ()=>{
    return(
        <div>Welcome User</div>
    )
}

GuestUser = ()=>{
    return(
        <div>Please login for better experience</div>
    );
}
}
export default ConditionalRendereing;