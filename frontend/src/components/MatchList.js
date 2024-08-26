import React, { useState, useEffect } from 'react';

const MatchList = () => {
    const [matches, setMatches] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        fetchMatches();
    }, []);

    const fetchMatches = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/matches?page=1&pageSize=10');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setMatches(data.matches);
            setLoading(false);
        } catch (error) {
            setError('Failed to fetch matches');
            setLoading(false);
        }
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;

    return (
        <div>
            <h2 className="text-xl font-bold mb-4">Upcoming Matches</h2>
            <ul>
                {matches.map((match) => (
                    <li key={match.ID} className="mb-2">
                        {match.Player1.Username} vs {match.Player2.Username} on {match.MatchDate}
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default MatchList;
