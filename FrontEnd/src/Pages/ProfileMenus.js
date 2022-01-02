import React, { useState } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import { View, Text, StyleSheet, ScrollView, Image, TouchableOpacity } from 'react-native';
// import { TouchableOpacity } from 'react-native-gesture-handler';
import { TextInput, Button, Gap } from '../components/atoms';
import { getData, storeData } from '../utils';
import { HeadProfile, ProfileTabs } from '../components';
import { BackIcon } from '../assets';

export default function ProfileMenus({ navigation }) {

    return (
        <ScrollView style={styles.wrapper}>
            <View style={styles.header}>
                <TouchableOpacity onPress={() => navigation.navigate('SignUp')}>
                    <Image source={BackIcon} style={styles.iconBack} />
                </TouchableOpacity>
            </View>
            <View style={{ flex: 1 }}>
                {/* <HeadProfile /> */}
                {/* <ProfileTabs /> */}
                {/* <MenusProfile /> */}
                <View
                    style={{
                        paddingHorizontal: 15,
                        paddingBottom: 10,
                    }}>
                    <TouchableOpacity style={styles.wrapperItem}>
                        <Text style={styles.textItem}>My e-Sertifikasi</Text>
                    </TouchableOpacity>
                    <Gap height={15} />
                    <View style={styles.wrapperItemInfo}>
                        <TouchableOpacity style={styles.itemInfo} onPress={() => navigation.navigate('AccountInformation')}>
                            <Text style={styles.textItem}>Account Information</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={styles.itemInfo}>
                            <Text style={styles.textItem}>Password & Security</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={styles.itemInfo}>
                            <Text style={styles.textItem}>Notification</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={styles.itemInfo}>
                            <Text style={styles.textItem}>Language</Text>
                        </TouchableOpacity>
                    </View>
                    <Gap height={20} />
                    <Text style={{ fontSize: 14, fontWeight: 'bold', color: '#2e3033' }}>
                        Tentang Aplikasi
                    </Text>
                    <Gap height={15} />
                    <View style={styles.wrapperItemInfo}>
                        <TouchableOpacity style={styles.itemInfo}>
                            <Text style={styles.textItem}>App Version</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={styles.itemInfo}>
                            <Text style={styles.textItem}>Terms & Conditions</Text>
                        </TouchableOpacity>
                        <TouchableOpacity style={styles.itemInfo}>
                            <Text style={styles.textItem}>Privacy Policy</Text>
                        </TouchableOpacity>
                    </View>
                    <Gap height={15} />
                    <View style={styles.wrapperItemInfo}>
                        <TouchableOpacity style={styles.itemInfo}>
                            <Text style={styles.textItem} onPress={() => navigation.navigate('SignUp')}>Log out</Text>
                        </TouchableOpacity>
                    </View>
                </View>
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
    iconBack: {
        width: 20,
        height: 20,
    },
    header: {
        height: 40,
        justifyContent: 'center',
    },
    container: {
        paddingHorizontal: 15,
        paddingVertical: 10,
        flex: 1,
    },
    wrapperItem: {
        width: '100%',
        height: 45,
        borderRadius: 8,
        backgroundColor: '#ffffff',
        borderStyle: 'solid',
        borderWidth: 1,
        borderColor: 'rgba(255, 255, 255, 0.08)',
        elevation: 1,
        padding: 10,
        justifyContent: 'center',
    },
    textItem: {
        fontFamily: 'TruenoLt',
        fontSize: 14,
        fontWeight: 'normal',
        color: '#22262f',
    },
    wrapperItemInfo: {
        width: '100%',
        // height: 230,
        borderStyle: 'solid',
        borderRadius: 8,
        backgroundColor: '#ffffff',
        borderColor: 'rgba(255, 255, 255, 0.08)',
        elevation: 4,
        paddingHorizontal: 10,
        justifyContent: 'flex-start',
    },
    itemInfo: {
        height: 230 / 4,
        justifyContent: 'center',
        borderBottomWidth: 1,
        borderBottomColor: '#f5f7fa',
        // backgroundColor: 'yellow',
    },
});
