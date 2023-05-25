import React, { useState, useEffect } from 'react';
import './App.css';
import AppRouter from './AppRouter';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  async function fetchUserData() {
    try {
      const response = await fetch('http://localhost:8080/validate', {
        method: 'GET',
        mode: 'cors',
        credentials: 'include',
      });

      if (response.status === 200) {
        setIsLoggedIn(true);
      }
    } catch (error) {
      console.error('Error fetching user data:', error);
    }
  }

  useEffect(() => {
    const cookieExists = document.cookie.includes('jwt=');
    console.log("checked cookie exists", cookieExists);

    setIsLoggedIn(cookieExists);

    if (cookieExists) {
      fetchUserData();
    }
  }, []);

  const [activeRoute, setActiveRoute] = useState(window.location.pathname);

  const handleLogout = async () => {
    try {
      const response = await fetch('http://localhost:8080/logout', {
        method: 'POST',
        mode: 'cors',
        credentials: 'include',
      });

      if (response.status === 200) {
        // Remove the "session" cookie and set isLoggedIn to false
        document.cookie = 'jwt=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;';
        setIsLoggedIn(false);
        // redirect to home
        setActiveRoute('/');
      } else {
        console.error('Error logging out:', response.statusText);
      }
    } catch (error) {
      console.error('Error logging out:', error);
    }
  };
  

  return (
    <div className="App">
      <header className="App-header">
        <ul id="nav-list">
          
              {isLoggedIn ? (
                <>
                  <li>
                    <a
                      className={activeRoute === '/users' ? 'active' : ''}
                      href="/users"
                      onClick={() => setActiveRoute('/users')}
                    >
                      Users
                    </a>
                  </li>
                  <li>
              <a
                className={activeRoute === '/upload' ? 'active' : ''}
                href="/upload"
                onClick={() => setActiveRoute('/upload')}
              >
                Upload
              </a>
            </li>
            <li>
            <a href="#" onClick={handleLogout}>
                Logout
              </a>
            </li>
                </>
              ) : (
                <>
                  <li>
                    <a
                      className={activeRoute === '/login' ? 'active' : ''}
                      href="/login"
                      onClick={() => setActiveRoute('/login')}
                    >
                      Login
                    </a>
                  </li>
                  <li>
                    <a
                      className={activeRoute === '/register' ? 'active' : ''}
                      href="/register"
                      onClick={() => setActiveRoute('/register')}
                    >
                      Register
                    </a>
                  </li>
                
                </>
              )} 
              
        </ul>
      </header>
      <AppRouter isLoggedIn={isLoggedIn} />
    </div>
  );
}

export default App;