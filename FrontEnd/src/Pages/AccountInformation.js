import React, { useEffect, useState } from "react";
import { View, Text, StyleSheet, ScrollView, Image, TouchableOpacity } from 'react-native';
import { BackIcon } from "../assets";
import { Gap } from '../components/atoms';
import { getData } from "../utils";
import { getEmail } from "../utils";
import axios from "axios";

export default function AccountInformation({ navigation }) {

    const [username, setUsername] = useState([]);
    const [name, setName] = useState([]);
    const [joinDate, setJoinDate] = useState([]);
    const [phone, setPhone] = useState([]);
    const [address, setAddress] = useState([]);
    const [identityType, setIdentityType] = useState([]);
    const [identityNo, setIdentityNo] = useState([]);
    const [emergencyCall, setEmergencyCall] = useState([]);

    const getData = async () => {
        const email = 'okkybagus87@gmail.com';
        try {
            const res = await axios.get("http://10.0.2.2:8080/get-user", {
                params: {
                    email
                }
            });
            setUsername(res.data.username);
            setName(res.data.name);
            setJoinDate(res.data.created_date);
            setPhone(res.data.phonenumber);
            setAddress(res.data.address_ktp);
            setIdentityNo(res.data.identity_no);
            setIdentityType(res.data.identity_type);
            setEmergencyCall(res.data.emergency_call);
        }
        catch (error) {
            alert(error.message);
        }
    };

    useEffect(() => {
        getData()
    }, [])

    return (
        <ScrollView style={styles.wrapper}>
            <View style={styles.header}>
                <TouchableOpacity onPress={() => navigation.navigate('SignUp')}>
                    <Image source={BackIcon} style={styles.iconBack} />
                </TouchableOpacity>
            </View>
            <View style={{ flex: 1 }}>
                <View style={styles.judul}>
                    <Text style={styles.textHeader}>Account Information</Text>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotak1}>
                        <Text style={styles.kontenKotak}>Username</Text>
                        <Text style={styles.dataUser}>{username}</Text>
                    </View>
                    <View style={styles.kotak}>
                        <Text style={styles.kontenKotak}>Name</Text>
                        <Text style={styles.dataUser}>{name}</Text>
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotak1}>
                        <Text style={styles.kontenKotak}>Join Date</Text>
                        <Text style={styles.dataUser}>{joinDate}</Text>
                    </View>
                    <View style={styles.kotak}>
                        <Text style={styles.kontenKotak}>Phone Number</Text>
                        <Text style={styles.dataUser}>{phone}</Text>
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Alamat</Text>
                        <Text style={styles.dataUser}>{address}</Text>
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Kode Pos</Text>
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Jenis Identitas</Text>
                        <Text style={styles.dataUser}>{identityType}</Text>
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Nomor Identitas</Text>
                        <Text style={styles.dataUser}>{identityNo}</Text>
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakAkhir}>
                        <Text style={styles.kontenKotak}>Emergency Call</Text>
                        <Text style={styles.dataUser}>{emergencyCall}</Text>
                    </View>
                </View>
            </View>
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    dataUser: {
        padding: 5,
        color: 'black',
    },
    kontenKotak: {
        fontSize: 16,
        padding: 5,
    },
    konten: {
        justifyContent: 'center',
        flexDirection: 'row',
    },
    wrapper: {
        paddingVertical: 10,
        flex: 1,
        backgroundColor: 'white',
    },
    kotak1: {
        justifyContent: 'center',
        width: '50%',
        borderRightWidth: 1,
        borderTopWidth: 1,
        borderColor: '#80979797',
        height: 80,
    },
    kotak: {
        justifyContent: 'center',
        width: '50%',
        borderTopWidth: 1,
        borderColor: '#80979797',
        height: 80,
    },
    kotakDua: {
        justifyContent: 'center',
        marginHorizontal: 10,
        width: '100%',
        borderTopWidth: 1,
        borderColor: '#80979797',
        height: 80,
    },
    kotakAkhir: {
        justifyContent: 'center',
        marginHorizontal: 10,
        width: '100%',
        borderTopWidth: 1,
        borderBottomWidth: 2,
        borderColor: '#80979797',
        height: 80,
    },
    header: {
        height: 40,
        justifyContent: 'center',
    },
    iconBack: {
        marginHorizontal: 15,
        width: 20,
        height: 20,
    },
    judul: {
        marginBottom: 20,
        alignItems: 'center',
        justifyContent: 'center',
    },
    textHeader: {
        fontSize: 18,
        color: 'black',
    },
});