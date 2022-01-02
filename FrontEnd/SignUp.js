import React, { Component, useState } from 'react';
import { View, Text, StyleSheet, ScrollView, Image, Button, TouchableOpacity, TouchableHighlight } from 'react-native';
import { RectButton } from 'react-native-gesture-handler';
import { TextInput, Buttons, Gap } from '../components/atoms';
import { BackIcon, PhotoDummy } from '../assets/Icons';
import { launchImageLibrary } from 'react-native-image-picker';
import { useDispatch } from 'react-redux';
import { simpanNama, useForm } from '../utils';
import FlashMessage, { showMessage } from 'react-native-flash-message';

export default function SignUp({ navigation }) {
    const [name, setName] = useState("");

    const onSubmit = () => {
        if (name == '') {
            showMessage({
                message: "Field nama harus diisi",
                type: "warning",
            });
        }
        else {
            simpanNama(name);
            navigation.navigate('SignUpForm');
        }
    };

    return (
        <ScrollView style={styles.wrapper}>
            <View style={styles.header}>
                <TouchableOpacity>
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
                        <Text style={styles.subtitleRight}>1/2</Text>
                    </View>
                    <Gap height={10} />
                    <Text style={styles.labelTitle}>Profile Photo</Text>
                    <View style={styles.photoWrapper}>
                        <TouchableOpacity style={styles.photoContainer}>
                            {/* {photo ? (
                                <Image source={photo} style={styles.photoContainer} />
                            ) : (
                                <PhotoDummy width="50" height="50" />
                            )} */}
                        </TouchableOpacity>
                        <View style={styles.addPhoto}>
                            <Text style={{ color: 'white' }}>+</Text>
                        </View>
                    </View>
                    <Gap height={10} />
                    <TextInput
                        label="Name"
                        placeholder="Enter your full name"
                        onChangeText={(value) => setName(value)}
                    />
                    <Gap height={40} />
                    <TouchableOpacity style={styles.btn} onPress={onSubmit}>
                        <Text style={styles.textBtn}>Next</Text>
                    </TouchableOpacity>
                    {/* <Buttons
                        text="Next"
                        textcolor='white'
                        backgroundcolor="#757575"
                        backgroundcoloronpress="#0c8eff"
                        onPress={onSubmit}
                    /> */}
                    <Gap height={70} />
                    <View style={styles.signInWrapper}>
                        <Text style={styles.textOuter}>Already a member?</Text>
                        <Text style={styles.textInner} onPress={() => navigation.navigate('ProfileMenus')}>
                            Sign In
                        </Text>
                    </View>
                    <Gap height={70} />
                </View>
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
        // backgroundColor: 'red',
        alignItems: 'center',
        width: 25,
        height: 25,
    },
    photoWrapper: {
        height: 90,
        width: 85,
        position: 'relative',
    },
    photoContainer: {
        height: 75,
        width: 75,
        borderWidth: 1,
        borderColor: '#80979797',
        borderRadius: 20,
        marginVertical: 10,
        justifyContent: 'center',
        alignItems: 'center',
    },
    addPhoto: {
        backgroundColor: '#0c8eff',
        height: 20,
        width: 20,
        borderRadius: 50,
        justifyContent: 'center',
        alignItems: 'center',
        position: 'absolute',
        bottom: 0,
        right: 5,
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
        // position: 'absolute',
        // bottom: 80,
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
});