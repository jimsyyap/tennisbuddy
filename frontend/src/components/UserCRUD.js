import React, { useState, useEffect } from 'react';
import axios from 'axios';

const UserList = () => {
    const [users, setUsers] = useState([]);

    useEffect(() => {
        fetchUsers();
    }, []);

    const fetchUsers = async () => {
        const response = await axios.get('http://localhost:8080/api/users');
        setUsers(response.data);
    };

    return (
        <div className="container mx-auto px-4">
            <h2 className="text-2xl font-bold mb-4">Users</h2>
            <ul className="space-y-2">
                {users.map(user => (
                    <li key={user.ID} className="bg-white shadow rounded-lg p-4">
                        {user.username} - {user.email}
                    </li>
                ))}
            </ul>
        </div>
    );
};

const UserForm = ({ user, onSubmit }) => {
    const [formData, setFormData] = useState(user || { username: '', email: '', password: '' });

    const handleChange = (e) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e) => {
        e.preventDefault();
        onSubmit(formData);
    };

    return (
        <form onSubmit={handleSubmit} className="space-y-4">
            <div>
                <label htmlFor="username" className="block text-sm font-medium text-gray-700">Username</label>
                <input
                    type="text"
                    name="username"
                    value={formData.username}
                    onChange={handleChange}
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>
            <div>
                <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email</label>
                <input
                    type="email"
                    name="email"
                    value={formData.email}
                    onChange={handleChange}
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>
            <div>
                <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
                <input
                    type="password"
                    name="password"
                    value={formData.password}
                    onChange={handleChange}
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
                />
            </div>
            <button type="submit" className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
                {user ? 'Update User' : 'Create User'}
            </button>
        </form>
    );
};

const UserCRUD = () => {
    const [users, setUsers] = useState([]);
    const [selectedUser, setSelectedUser] = useState(null);

    useEffect(() => {
        fetchUsers();
    }, []);

    const fetchUsers = async () => {
        const response = await axios.get('http://localhost:8080/api/users');
        setUsers(response.data);
    };

    const createUser = async (userData) => {
        await axios.post('http://localhost:8080/api/users', userData);
        fetchUsers();
    };

    const updateUser = async (userData) => {
        await axios.put(`http://localhost:8080/api/users/${userData.ID}`, userData);
        fetchUsers();
        setSelectedUser(null);
    };

    const deleteUser = async (userId) => {
        await axios.delete(`http://localhost:8080/api/users/${userId}`);
        fetchUsers();
    };

    return (
        <div className="container mx-auto px-4">
            <h1 className="text-3xl font-bold mb-8">User Management</h1>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div>
                    <h2 className="text-2xl font-bold mb-4">User List</h2>
                    <ul className="space-y-2">
                        {users.map(user => (
                            <li key={user.ID} className="bg-white shadow rounded-lg p-4 flex justify-between items-center">
                                <span>{user.username} - {user.email}</span>
                                <div>
                                    <button onClick={() => setSelectedUser(user)} className="text-blue-500 hover:text-blue-700 mr-2">Edit</button>
                                    <button onClick={() => deleteUser(user.ID)} className="text-red-500 hover:text-red-700">Delete</button>
                                </div>
                            </li>
                        ))}
                    </ul>
                </div>
                <div>
                    <h2 className="text-2xl font-bold mb-4">{selectedUser ? 'Edit User' : 'Create User'}</h2>
                    <UserForm
                        user={selectedUser}
                        onSubmit={selectedUser ? updateUser : createUser}
                    />
                </div>
            </div>
        </div>
    );
};

export default UserCRUD;
