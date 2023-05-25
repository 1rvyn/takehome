import React, { useState, useEffect } from 'react';
import { AiFillDelete, AiFillEdit } from 'react-icons/ai';
import Modal from 'react-modal';
import axios from 'axios';
import './App.css';

Modal.setAppElement('#root');

const customStyles = {
  content: {
    top: '50%',
    left: '50%',
    right: 'auto',
    bottom: 'auto',
    marginRight: '-50%',
    transform: 'translate(-50%, -50%)',
    backgroundColor: '#888',
  },
  overlay: {
    backgroundColor: 'rgba(0, 0, 0, 0.75)'
  }
};

function Users(props) {
  const { isLoggedIn } = props;
  const [users, setUsers] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const usersPerPage = 10;
  const [currentUser, setCurrentUser] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [firstName, setFirstName] = useState('');
  const [surname, setSurname] = useState('');
  const [email, setEmail] = useState('');
  const [dob, setDob] = useState('');
  const [filePath, setFilePath] = useState('');
  const [isConfirmationOpen, setIsConfirmationOpen] = useState(false);
  const [userToDelete, setUserToDelete] = useState(null);
  const [isNewUser, setIsNewUser] = useState(false); 
  const [password, setPassword] = useState(''); 



  useEffect(() => {
    const getUsers = async () => {
      const res = await axios.get('http://localhost:8080/users', { withCredentials: true });
      const usersArray = Object.values(res.data);
      setUsers(usersArray);
    };
    getUsers();
  }, []);
  

  const indexOfLastUser = currentPage * usersPerPage;
  const indexOfFirstUser = indexOfLastUser - usersPerPage;
  const currentUsers = users.slice(indexOfFirstUser, Math.min(indexOfLastUser, users.length));

  const nextPage = () => {
    if (currentUsers.length === usersPerPage) {
      setCurrentPage(currentPage + 1);
    }
  };

  const prevPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const deleteUser = async (userId) => {
    await axios.delete(`http://localhost:8080/user/${userId}`, { withCredentials: true });
    const res = await axios.get('http://localhost:8080/users', { withCredentials: true });
    const usersArray = Object.values(res.data);
    setUsers(usersArray);
  };

  const openEditModal = (user) => {
    setIsNewUser(false); 
    setFirstName(user.FirstName);
    setSurname(user.Surname);
    setEmail(user.Email);
    setDob(user.DOB);
    setFilePath(user.FilePath);
    setCurrentUser(user);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setIsConfirmationOpen(false);
    setUserToDelete(null);
  };

  const updateUser = async () => {
    if (isNewUser) {
      await axios.post(`http://localhost:8080/create`, {
        FirstName: firstName,
        Surname: surname,
        Email: email,
        DOB: dob,
        Password: password,
        FilePath: filePath
      }, { withCredentials: true });
    } else {
      await axios.put(`http://localhost:8080/update/${currentUser.ID}`, {
        FirstName: firstName,
        Surname: surname,
        Email: email,
        DOB: dob,
        FilePath: filePath
      }, { withCredentials: true });
    }
    
    closeModal();
    // refresh users
    const res = await axios.get('http://localhost:8080/users', { withCredentials: true });
    const usersArray = Object.values(res.data);
    setUsers(usersArray);
  };

  const createUser = () => {
    setIsNewUser(true);  
    setCurrentUser(null);  
    setFirstName('');
    setSurname('');
    setPassword('');
    setEmail('');
    setDob('');
    setFilePath('');
    setIsModalOpen(true);
  };

  const confirmDeleteUser = (user) => {
    setUserToDelete(user);
    setIsConfirmationOpen(true);
  };

  const handleDeleteUser = () => {
    deleteUser(userToDelete.ID);
    closeModal();
  };

  return (
    <div>
      {isLoggedIn ? (
        <>
          <br />
          <p>You are logged in. Show users!</p>
          <button onClick={prevPage} disabled={currentPage === 1}>
            Previous Page
          </button>
          <button onClick={nextPage} disabled={currentUsers.length < usersPerPage}>
            Next Page
          </button>
          <p>Page: {currentPage}</p>
          <ul>
            {currentUsers.map((user) => (
              <div key={user.ID} style={{display: 'flex', alignItems: 'center'}}>
                <li style={{marginRight: '1rem'}}>
                  {user.FirstName} {user.Surname} ({user.Email}, {user.DOB}, {user.FilePath || "No file path available"})
                </li>
                <button onClick={() => confirmDeleteUser(user)} style={{marginRight: '0.5rem'}}>
                  <AiFillDelete />
                </button>
                <button onClick={() => openEditModal(user)}>
                  <AiFillEdit />
                </button>
              </div>
            ))}
          </ul>
          <button onClick={createUser}>Create User</button>
          <Modal
              isOpen={isModalOpen}
              onRequestClose={closeModal}
              contentLabel={isNewUser ? "Create User" : "Edit User"}
              style={customStyles}
            >
               <h2>{isNewUser ? "Create User" : "Edit User"}</h2>
              
            <h2>Edit User</h2>
            <label>First Name: <input type="text" value={firstName} onChange={e => setFirstName(e.target.value)} /></label>
            <label>Surname: <input type="text" value={surname} onChange={e => setSurname(e.target.value)} /></label>
            <label>Email: <input type="text" value={email} onChange={e => setEmail(e.target.value)} /></label>
            <label>Date of Birth: <input type="text" value={dob} onChange={e => setDob(e.target.value)} /></label>
  
            {isNewUser && (
              <label>Password: <input type="password" value={password} onChange={e => setPassword(e.target.value)} /></label>
            )}

            <button onClick={updateUser}>Confirm</button>
            <button onClick={closeModal}>Cancel</button>

            
          </Modal>
          <Modal
            isOpen={isConfirmationOpen}
            onRequestClose={closeModal}
            contentLabel="Delete User Confirmation"
            style={customStyles}
          >
            <h2>Are you sure you want to delete this user?</h2>
            <p>{userToDelete && `${userToDelete.FirstName} ${userToDelete.Surname} (${userToDelete.Email}, ${userToDelete.DOB}, ${userToDelete.FilePath || "No file path available"})`}</p>
            <button onClick={handleDeleteUser}>Yes</button>
            <button onClick={closeModal}>No</button>
          </Modal>
          
        </>

      ) : (
        <p>Please log in to view users.</p>
      )}
    </div>
  );
}

export default Users;