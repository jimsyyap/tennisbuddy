import React, { useState, useEffect } from 'react';
import axios from 'axios';

const App = () => {
    const [message, setMessage] = useState('');
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get('http://localhost:8080/api/hello');
                setMessage(response.data.message);
            } catch (err) {
                setError('Failed to connect to the server');
            }
        };

        fetchData();
    }, []);

    return (
        <div className="min-h-screen bg-gray-100 flex items-center justify-center px-4">
            <div className="max-w-md w-full bg-white rounded-lg shadow-md p-8">
                <h1 className="text-2xl font-bold text-center mb-4">TennisBuddy</h1>
                {error ? (
                    <p className="text-red-500 text-center">{error}</p>
                ) : (
                    <p className="text-gray-700 text-center">{message}</p>
                )}
            </div>
        </div>
    );
};

export default App;
