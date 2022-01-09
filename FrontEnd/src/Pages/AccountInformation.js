import React, { useEffect, useState } from "react";
import { View, Text, StyleSheet, ScrollView, Image, TouchableOpacity, TextInput, SafeAreaView } from 'react-native';
import { BackIcon } from "../assets";
import { Gap } from '../components/atoms';
import { getData } from "../utils";
import { getEmail } from "../utils";
import axios from "axios";
// import { TextInput } from "react-native-gesture-handler";
import { showMessage } from "react-native-flash-message";
import { Dropdown, DropDown } from "react-native-element-dropdown";

const email = 'okkybagus87@gmail.com';

export default function AccountInformation({ navigation }) {

    const dd = [
        { label: 'KTP', value: 'KTP' },
        { label: 'SIM', value: 'SIM' },
        { label: 'KK', value: 'KK' },
    ];
    const [open, setOpen] = useState(false);
    const [username, setUsername] = useState('');
    const [name, setName] = useState('');
    const [joinDate, setJoinDate] = useState('');
    const [phone, setPhone] = useState('');
    const [address, setAddress] = useState('');
    const [identityType, setIdentityType] = useState(null);
    const [identityNo, setIdentityNo] = useState('');
    const [emergencyCall, setEmergencyCall] = useState('');
    const [kodePos, setKodePos] = useState('');

    const getData = async () => {
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
            setKodePos(res.data.postal_code);
        }
        catch (error) {
            alert(error.message);
        }
    }

    const dropDown = () => {

    }

    const onSubmit = async () => {
        try {
            const data = {
                email: email,
                username: username,
                name: name,
                phonenumber: phone,
                address_ktp: address,
                postal_code: kodePos,
                identity_type: identityType,
                identity_no: identityNo,
                emergency_call: emergencyCall,
            }
            const res = await axios.post('http://10.0.2.2:8080/update-profile', data);
            if ((await res).data.status == true) {
                showMessage({
                    message: "Informasi akun berhasil diubah",
                    type: 'info',
                });
            }
            else {
                showMessage({
                    message: "Gagal",
                    type: 'warning',
                });
            }
        } catch (error) {
            alert(error.message);
        }
    }

    useEffect(() => {
        getData()
    }, [])

    return (
        <ScrollView style={styles.wrapper} nestedScrollEnabled={true}>
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
                        {/* <Text style={styles.dataUser}>{username}</Text> */}
                        <TextInput
                            autoCapitalize="none"
                            placeholder="Enter your username"
                            value={username}
                            onChangeText={(value) => setUsername(value)}
                        />
                    </View>
                    <View style={styles.kotak}>
                        <Text style={styles.kontenKotak}>Name</Text>
                        {/* <Text style={styles.dataUser}>{name}</Text> */}
                        <TextInput
                            autoCapitalize="none"
                            placeholder="Enter your name"
                            value={name}
                            onChangeText={(value) => setName(value)}
                        />
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotak1}>
                        <Text style={styles.kontenKotak}>Join Date</Text>
                        <Text style={styles.dataUser}>{joinDate}</Text>
                    </View>
                    <View style={styles.kotak}>
                        <Text style={styles.kontenKotak}>Phone Number</Text>
                        <TextInput
                            keyboardType="numeric"
                            autoCapitalize="none"
                            placeholder="Enter your phone number"
                            value={phone}
                            onChangeText={(value) => setPhone(value)}
                        />
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Alamat</Text>
                        {/* <Text style={styles.dataUser}>{address}</Text> */}
                        <TextInput
                            autoCapitalize="none"
                            placeholder="Enter your address"
                            value={address}
                            onChangeText={(value) => setAddress(value)}
                        />
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Kode Pos</Text>
                        <TextInput
                            autoCapitalize="none"
                            placeholder="Enter your postal code"
                            value={kodePos}
                            onChangeText={(value) => setKodePos(value)}
                        />
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Jenis Identitas</Text>
                        <Dropdown
                            style={styles.dropdown}
                            placeholderStyle={styles.placeholderStyle}
                            selectedTextStyle={styles.selectedTextStyle}
                            inputSearchStyle={styles.inputSearchStyle}
                            iconStyle={styles.iconStyle}
                            data={dd}
                            maxHeight={169}
                            labelField="label"
                            valueField="value"
                            placeholder="Select Item"
                            value={identityType}
                            onChange={item=>{
                                setIdentityType(item.value);
                            }}
                        />
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakDua}>
                        <Text style={styles.kontenKotak}>Nomor Identitas</Text>
                        <TextInput
                            autoCapitalize="none"
                            placeholder="Enter your identity number"
                            value={identityNo}
                            onChangeText={(value) => setIdentityNo(value)}
                        />
                    </View>
                </View>
                <View style={styles.konten}>
                    <View style={styles.kotakAkhir}>
                        <Text style={styles.kontenKotak}>Emergency Call</Text>
                        <TextInput
                            autoCapitalize="none"
                            placeholder="Enter your emergency number"
                            value={emergencyCall}
                            onChangeText={(value) => setEmergencyCall(value)}
                        />
                    </View>
                </View>
                <TouchableOpacity style={styles.btn} onPress={onSubmit}>
                    <Text style={styles.textBtn}>Save</Text>
                </TouchableOpacity>
            </View>
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    dropdown: {
        marginHorizontal: 10,
        height: 40,
    },
    icon: {
        marginRight: 5,
    },
    placeholderStyle: {
        fontSize: 16,
    },
    selectedTextStyle: {
        fontSize: 16,
    },
    iconStyle: {
        width: 20,
        height: 20,
    },
    inputSearchStyle: {
        height: 40,
        fontSize: 16,
    },
    textBtn: {
        color: 'white',
        fontSize: 17,
        fontFamily: 'RobotoMedium',
        textAlign: 'center',
    },
    btn: {
        margin: 20,
        borderRadius: 20,
        borderWidth: 1,
        backgroundColor: '#0c8eff',
        borderColor: '#80979797',
        justifyContent: 'center',
        height: 60,
    },
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