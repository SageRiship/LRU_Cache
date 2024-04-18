import React, { useState } from 'react';
import axios from 'axios';

function App() {
  const [key, setKey] = useState('');
  const [value, setValue] = useState('');
  const [expiration, setExpiration] = useState('');
  const [getResponse, setGetResponse] = useState('');
  const [setResponse, setSetResponse] = useState('');

  const handleGet = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/cache/${key}`);
      setGetResponse(response.data);
    } catch (error) {
      console.error('Error fetching data:', error);
    }
  };

  const handleSet = async () => {
    try {
      const response = await axios.post('http://localhost:8080/cache', {
        key,
        value,
        expiration: parseInt(expiration),
      });
      setSetResponse('Value set successfully');
    } catch (error) {
      console.error('Error setting data:', error);
    }
  };

  return (
    <div>
      <h1>LRU Cache React App - Apica</h1>
      <div>
        <h2>Get Value from Cache</h2>
        <input type="text" placeholder="Enter Key" value={key} onChange={(e) => setKey(e.target.value)} />
        <button onClick={handleGet}>Get Value</button>
        <p>Response: {getResponse}</p>
      </div>
      <div>
        <h2>Set Value in Cache</h2>
        <input type="text" placeholder="Enter Key" value={key} onChange={(e) => setKey(e.target.value)} />
        <input type="text" placeholder="Enter Value" value={value} onChange={(e) => setValue(e.target.value)} />
        <input type="text" placeholder="Enter Expiration (seconds)" value={expiration} onChange={(e) => setExpiration(e.target.value)} />
        <button onClick={handleSet}>Set Value</button>
        <p>{setResponse}</p>
      </div>
    </div>
  );
}

export default App;
