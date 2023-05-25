import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';

import Home from './Home';
import Login from './Login';
import Register from './Register';
import Users from './Users';
import Upload from './Upload';

function AppRouter({ isLoggedIn, isAdmin }) {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home isLoggedIn={isLoggedIn} />} />

        <Route path="/Users" element={<Users isLoggedIn={isLoggedIn} />} />


        <Route path="/Upload" element={<Upload isLoggedIn={isLoggedIn} />} />

        <Route path="/login" element={<Login />} />

        <Route path="/register" element={<Register />} />
      </Routes>
    </BrowserRouter>
  );
}

export default AppRouter;
