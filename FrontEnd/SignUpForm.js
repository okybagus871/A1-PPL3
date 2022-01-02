import React, { Component, useState } from 'react'
import { View, Text, StyleSheet, ScrollView, Image, Button, TouchableOpacity } from 'react-native';
// import { TouchableOpacity } from 'react-native-gesture-handler';
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
import { getEmail, simpanKeDb, useForm, simpanEmail, getNama } from '../utils';
import axios from 'axios';
import { format } from 'jest-validate';
import FlashMessage, { showMessage } from 'react-native-flash-message';

export default function SignUpForm({ navigation }) {
    const [email, setEmail] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [confirmPassword, setConfirmPassword] = useState("");

    const checkUsername = async (username) => {
        try {
            const res = await axios.get('http://10.0.2.2:8080/check-username', {
                params: {
                    username
                }
            });
            if (res.data.status == true) {
                setUsername(username);
            } else {
                showMessage({
                    message: "Username " + username + " has already taken",
                    type: "warning",
                });
            }
        } catch (error) {
            alert(error.message);
        }
    }

    const checkEmail = async (email) => {
        try {
            const res = await axios.get('http://10.0.2.2:8080/check-email', {
                params: {
                    email
                }
            });
            if (res.data.status == true) {
                setEmail(email);
            } else {
                showMessage({
                    message: "Email " + email + " has already taken",
                    type: "warning",
                });
            }
        } catch (error) {
            alert(error.message);
        }
    }

    const onSubmit = async () => {
        // try {
        //     const respon = await axios.get('https://api.kickbox.com/v2/verify', {
        //         params: {
        //             email,
        //             apikey: 'live_be3b8dcceea5a956bc10a26e6068f9001925aef9b5f3cab17c195e2eefe362a4'
        //         }
        //     });
        //     if (respon.data.result == 'undeliverable'){
        //         showMessage({
        //             message: "Invalid email",
        //             type: "warning",
        //         });
        //     }
        //     else if (confirmPassword != password) {
        //         showMessage({
        //             message: "Password not match",
        //             type: "warning",
        //         });
        //     }
        //     else {
        //         showMessage({
        //             message: "Check your email to verification account",
        //             type: "info",
        //         });
        //         simpanEmail(email);
        //         const name = getNama();
        //         const data = {
        //             name,
        //             email,
        //             username,
        //             password,
        //         };
        //         axios.post("http://10.0.2.2:8080/signup", data);
        //         navigation.navigate('CheckEmailToken');
        //     }
        // } catch (error) {
        //     alert(error.message);
        // }
        if (confirmPassword != password) {
            showMessage({
                message: "Password not match",
                type: "warning",
            });
        }
        else {
            simpanEmail(email);
            const name = getNama();
            const data = {
                name,
                email,
                username,
                password,
            };
            axios.post("http://10.0.2.2:8080/signup", data);
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
                        onChangeText={(value) => checkUsername(value)}
                    />
                    <Gap height={20} />
                    <TextInput
                        label="Email"
                        placeholder="Enter your email"
                        textContentType="emailAddress"
                        keyboardType="email-address"
                        autoCapitalize="none"
                        onChangeText={(value) => checkEmail(value)}
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
                    <TouchableOpacity style={styles.btn} onPress={onSubmit} disabled={username && email && password && confirmPassword ? false : true}>
                        <Text style={styles.textBtn}>Create Account</Text>
                    </TouchableOpacity>
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
    textBtn: {
        color: 'white',
        fontSize: 14,
        fontFamily: 'RobotoMedium',
        textAlign: 'center',
    },
    btn: {
        borderRadius: 20,
        borderWidth: 1,
        backgroundColor: '#0c8eff',
        borderColor: '#80979797',
        justifyContent: 'center',
        height: 60,
    },
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