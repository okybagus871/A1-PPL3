import axios from 'axios';
import useForm from './useForm';
export * from './showMessage';
export * from './storage';

var name;
var email;

export const simpanNama = (nama) => {
    name = nama;
};

export const getNama = () => {
    return name;
};

export const simpanEmail = (emal) => {
    email = emal;
};

export const getEmail = () => {
    return email;
};

export {useForm};
