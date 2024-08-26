import React, { useState } from 'react';
import { Formik, Form, Field, ErrorMessage } from 'formik';
import * as Yup from 'yup';
import axios from 'axios';
import './UserRegistrationForm.css';

const validationSchema = Yup.object().shape({
    username: Yup.string()
        .min(3, 'Username must be at least 3 characters')
        .required('Username is required'),
    email: Yup.string()
        .email('Invalid email address')
        .required('Email is required'),
    password: Yup.string()
        .min(8, 'Password must be at least 8 characters')
        .required('Password is required'),
});

const UserRegistrationForm = () => {
    const [serverError, setServerError] = useState('');

    const initialValues = {
        username: '',
        email: '',
        password: '',
    };

    const handleSubmit = async (values, { setSubmitting, resetForm }) => {
        try {
            const response = await axios.post('/api/users', values);
            console.log('User registered successfully:', response.data);
            resetForm();
            // Handle successful registration (e.g., show success message, redirect)
        } catch (error) {
            setServerError(error.response?.data?.error || 'Failed to register user');
        } finally {
            setSubmitting(false);
        }
    };

    return (
        <div className="form-container">
            <h2>User Registration</h2>
            <Formik
                initialValues={initialValues}
                validationSchema={validationSchema}
                onSubmit={handleSubmit}
            >
                {({ isSubmitting }) => (
                    <Form>
                        <div className="form-field">
                            <label htmlFor="username">Username:</label>
                            <Field type="text" name="username" />
                            <ErrorMessage name="username" component="div" className="error" />
                        </div>

                        <div className="form-field">
                            <label htmlFor="email">Email:</label>
                            <Field type="email" name="email" />
                            <ErrorMessage name="email" component="div" className="error" />
                        </div>

                        <div className="form-field">
                            <label htmlFor="password">Password:</label>
                            <Field type="password" name="password" />
                            <ErrorMessage name="password" component="div" className="error" />
                        </div>

                        <button type="submit" disabled={isSubmitting}>
                            Register
                        </button>

                        {serverError && <div className="error">{serverError}</div>}
                    </Form>
                )}
            </Formik>
        </div>
    );
};

export default UserRegistrationForm;
