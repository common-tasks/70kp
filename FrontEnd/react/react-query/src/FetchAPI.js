import axios from "axios";

async function fetchPosts() {
    const {data} = await axios.get('http://localhost:9999/health-check');
    console.log('data received from rest service ' + data);
    return data;
}

export default fetchPosts;