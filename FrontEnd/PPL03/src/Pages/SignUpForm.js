import React, { Component, useState } from 'react'
import { View, Text, StyleSheet, ScrollView, Image, Button } from 'react-native';
import { TouchableOpacity } from 'react-native-gesture-handler';
import {
    TextInputPassword,
    TextInput,
    Buttons,
    Gap,
} from '../components/atoms';
import {
    BackIcon,
    LockIcon,
    SeePassword,
    LockIconActive,
    SeePasswordActive,
} from '../assets';
import { useDispatch, useSelector } from 'react-redux';
import { signUpAction } from '../redux/action';
import { getEmail, simpanKeDb, useForm } from '../utils';
import axios from 'axios';
import { format } from 'jest-validate';
import FlashMessage, { showMessage } from 'react-native-flash-message';

export default function SignUpForm({ navigation }) {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    // const [form, setForm] = useForm({
    //     email: '',
    //     username: '',
    //     password: '',
    // })

    // const param = JSON.stringify({
    //     "email": value.email,
    //     "password": value.password,
    //     "name": value.name,
    //     "username": value.username,
    // });

    const onSubmit = () => {
        if (username == '') {
            showMessage({
                message: "Isi semua field",
                type: "warning",
            });
        }
        else if (confirmPassword != password){
            showMessage({
                message: "Password tidak sama",
                type: "warning",
            });
        }
        else {
            showMessage({
                message: "Kode verifikasi berhasil dikirim ke email anda",
                type: "info",
            });
            simpanKeDb(email, username, password);
            navigation.navigate('CheckEmailToken');
        }
    }

    return (
        <ScrollView style={styles.wrapper}>
            <View style={styles.header}>
                <TouchableOpacity onPress={() => navigation.navigate('SignUp')}>
                    <Image source={BackIcon} style={styles.iconBack}></Image>
                </TouchableOpacity>
            </View>
            <View style={styles.container}>
                <View style={styles.main}>
                    <Text style={styles.titleText}>Welcome to sadhelX</Text>
                    <View style={styles.subtitleContainer}>
                        <Text style={styles.subtitleLeft}>
                            We just need a few quick details to continue
                        </Text>
                        <Text style={styles.subtitleRight}>2/2</Text>
                    </View>
                    <Gap height={20} />
                    <TextInput
                        label="Username"
                        autoCapitalize="none"
                        placeholder="Enter your username"
                        onChangeText={(value) => setUsername(value)}
                    />
                    <Gap height={20} />
                    <TextInput
                        label="Email"
                        placeholder="Enter your email"
                        textContentType="emailAddress"
                        keyboardType="email-address"
                        autoCapitalize="none"
                        onChangeText={(value) => setEmail(value)}
                    />
                    <Gap height={20} />
                    <TextInputPassword
                        label="Create Password"
                        placeholder="Password at least 8 characters"
                        sourceImageLeft={LockIcon}
                        sourceImageLeftActive={LockIconActive}
                        sourceImageRight={SeePassword}
                        sourceImageRightActive={SeePasswordActive}
                        onChangeText={(value) => setPassword(value)}
                    />
                    <Gap height={20} />
                    <TextInputPassword
                        label="Confirm Password"
                        placeholder="Password at least 8 characters"
                        sourceImageLeft={LockIcon}
                        sourceImageLeftActive={LockIconActive}
                        sourceImageRight={SeePassword}
                        sourceImageRightActive={SeePasswordActive}
                        onChangeText={(value) => setConfirmPassword(value)}
                    />
                    <Gap height={10} />
                    <Text style={styles.confrimtext}>Both password must be match</Text>
                    <Gap height={20} />
                    <Buttons
                        backgroundcolor="#0c8eff"
                        backgroundcoloronpress="#0c8eff"
                        textcolor="#ffff"
                        text="Create Account"
                        onPress={onSubmit}
                    />
                    <Gap height={30} />
                    <View style={styles.signInWrapper}>
                        <Text style={styles.textOuter}>Already a member?</Text>
                        <Text style={styles.textInner} onPress={() => handleGoTo('SignIn')}>
                            Sign In
                        </Text>
                    </View>
                </View>
                <Gap height={30} />
            </View>
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    wrapper: {
        paddingHorizontal: 15,
        paddingVertical: 10,
        flex: 1,
        backgroundColor: 'white',
    },
    header: {
        height: 40,
        justifyContent: 'center',
    },
    container: {
        flex: 1,
    },
    main: {
        paddingVertical: 10,
        paddingHorizontal: 25,
        flex: 1,
    },
    iconBack: {
        width: 20,
        height: 20,
    },
    titleText: {
        color: 'black',
        fontSize: 26,
        fontFamily: 'SarabunExtraBold',
    },
    subtitleContainer: {
        flexDirection: 'row',
        width: '100%',
        justifyContent: 'space-between',
    },
    subtitleLeft: {
        fontFamily: 'SarabunRegular',
        width: 200,
        color: '#6e6e6e',
    },
    subtitleRight: {
        fontFamily: 'SarabunMedium',
        fontSize: 16,

        alignItems: 'center',
        width: 25,
        height: 25,
    },
    labelTitle: {
        fontFamily: 'RobotoRegular',
        fontSize: 12,
        color: '#22262f',
    },
    signInWrapper: {
        flexDirection: 'row',
        width: '100%',
        justifyContent: 'center',
        fontSize: 14,
    },
    textOuter: {
        color: '#22262f',
        fontFamily: 'SarabunMedium',
    },
    textInner: {
        fontFamily: 'SarabunMedium',
        marginLeft: 5,
        color: '#0c8eff',
    },
    confrimtext: {
        color: '#6e6e6e',
        fontSize: 12,
    },
});