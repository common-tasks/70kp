import axios from "axios";

async function getHealth() {
    const {data} = await axios.get('http://localhost:9999/health-check');
    console.log('data received from rest service ' + data);
    return data;
}
async function getServices() {
    const {data} = await axios.get('http://localhost:9999/allServices');
    console.log('data received from allServices ' + JSON.stringify(data));
    return data;
}
export {getHealth,getServices};