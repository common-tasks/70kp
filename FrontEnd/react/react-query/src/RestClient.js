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
async function getService(serviceName) {
    const {data} = await axios.get(`http://localhost:9999/service/${serviceName}`);
    console.log('data received from specific service request ' +  JSON.stringify(data));
    return data;
}

async function updateService(service) {

    const {data} = await axios.post('http://localhost:9999/services',{
        name: 'calculator',
        id: 'xyz',
      });
      console.log('post service data'+ JSON.stringify(data));
      return data;
}
export {getHealth,getServices,getService,updateService};