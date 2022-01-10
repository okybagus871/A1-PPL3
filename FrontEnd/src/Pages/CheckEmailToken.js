import React, { useState } from 'react';
import {
    View,
    Text,
    StyleSheet,
    ScrollView,
    Image,
    TextInput,
    TouchableOpacity,
} from 'react-native';
// import { TouchableOpacity } from 'react-native-gesture-handler';
import { Buttons, Gap } from '../components/atoms';
import { BackIcon } from '../assets';
import SmoothPinCodeInput from 'react-native-smooth-pincode-input';
import { useDispatch, useSelector } from 'react-redux';
import { getEmail, simpanKeDb, useForm, verifyEmail } from '../utils';
import { verifyAction } from '../redux/action';
import FlashMessage, { showMessage } from 'react-native-flash-message';
import axios from 'axios';

export default function CheckEmailToken({ navigation }) {
    const [otp, setOtp] = useState('');

    const onSubmit = async () => {
        const email = getEmail();
        var otpInt = parseInt(otp);
        const data = {
            email,
            otp: otpInt,
        };
        try {
            const res = axios.post("http://10.0.2.2:8080/verify/mail", data);
            if ((await res).data.status == false) {
                showMessage({
                    message: "Invalid OTP",
                    type: 'warning',
                });
            }
            else{
                navigation.navigate('SuccessSignUp');
            }
        } catch (error) {
            alert(error.message);
        }
    }

    return (
        <ScrollView style={styles.wrapper}>
            <View style={styles.header}>
                <TouchableOpacity onPress={() => handleGoTo('SignUp')}>
                    <Image source={BackIcon} style={styles.iconBack} />
                </TouchableOpacity>
            </View>
            <View style={styles.container}>
                <View style={styles.main}>
                    <View style={styles.title}>
                        <Text style={styles.titleText}>Check your email</Text>
                        <View style={styles.subtitleContainer}>
                            <Text style={styles.subtitle}>
                                We have sent a verify account {'\n'}instruction to your email
                            </Text>
                            <Gap height={10} />
                        </View>
                    </View>
                    <Gap height={10} />
                    {/* <View style={styles.verifyContainer}> */}
                    <SmoothPinCodeInput
                        keyboardType="numeric"
                        containerStyle={styles.verifyContainer}
                        codeLength={4}
                        cellStyleFocused={styles.cellFocused}
                        cellStyle={styles.numberInput}
                        value={otp}
                        onTextChange={(pin) => setOtp(pin)}
                    />
                    {/* </View> */}
                    <Gap height={20} />
                    <Text style={styles.subtitle}>
                        Did not receive the email? Check your spam filter, or{' '}
                        <Text style={styles.subtitleEmail}> try another email address</Text>
                    </Text>
                    <Gap height={40} />
                    <Gap height={40} />
                    <TouchableOpacity style={styles.btn} onPress={onSubmit} disabled={otp ? false : true}>
                        <Text style={styles.textBtn}>Verify Email</Text>
                    </TouchableOpacity>
                    <Gap height={100} />
                    <View style={styles.helpWrapper}>
                        <Text
                            style={styles.textInner}
                            onPress={() => handleGoTo('WelcomeAuth')}>
                            Need more help?
                        </Text>
                    </View>
                    <Gap height={100} />
                </View>
            </View>
        </ScrollView>
    );
}

const styles = StyleSheet.create({
    textBtn: {
        color: 'white',
        fontSize: 17,
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
        fontSize: 25,
        fontFamily: 'SarabunExtraBold',
        textAlign: 'center',
    },
    title: {
        paddingHorizontal: 25,
    },
    subtitleContainer: {
        flexDirection: 'column',
        width: '100%',
    },
    subtitle: {
        fontFamily: 'SarabunRegular',
        fontSize: 15,
        textAlign: 'center',
        color: '#6e6e6e',
    },
    subtitleEmail: {
        marginTop: -8,
        fontFamily: 'SarabunRegular',
        fontSize: 15,
        textAlign: 'center',
        color: '#0c8eff',
    },
    verifyContainer: {
        height: 80,
        width: '100%',
    },
    numberInput: {
        height: 50,
        borderWidth: 1,
        borderColor: '#80979797',
        borderRadius: 10,
        padding: 10,
        marginRight: 5,
    },
    helpWrapper: {
        flexDirection: 'row',
        width: '100%',
        justifyContent: 'center',
        fontSize: 14,
    },
    textInner: {
        fontFamily: 'SarabunMedium',
        marginLeft: 5,
        color: '#0c8eff',
    },
    cellFocused: {
        borderColor: '#0c8eff',
        borderWidth: 2,
    },
});