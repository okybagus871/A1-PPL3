import axios from 'axios';
import useForm from './useForm';
export * from './showMessage';
export * from './storage';

var name;

export const simpanNama = (nama) => {
    name = nama;
};

export const simpanKeDb = (email, username, password) => {
    email = email;
    const data = {
        name,
        email,
        username,
        password,
    };
    axios.post("http://10.0.2.2:8080/signup", data);
};

export {useForm};
