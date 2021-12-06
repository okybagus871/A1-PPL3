// In App.js in a new project

import * as React from 'react';
import { View, Text } from 'react-native';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import SignUp from './Pages/SignUp';
import SignUpForm from './Pages/SignUpForm';
import CheckEmailToken from './Pages/CheckEmailToken';
import SuccessSignUp from './Pages/SuccessSignUp';
import FlashMessage from 'react-native-flash-message';

const Stack = createNativeStackNavigator();

function App() {
    return (
        <NavigationContainer>
            <Stack.Navigator screenOptions={{ headerShown: false }}>
                <Stack.Screen name="SignUp" component={SignUp} />
                <Stack.Screen name="SignUpForm" component={SignUpForm} />
                <Stack.Screen name="CheckEmailToken" component={CheckEmailToken} />
                <Stack.Screen name="SuccessSignUp" component={SuccessSignUp} />
            </Stack.Navigator>
            <FlashMessage style={{
                borderRadius: 10,
                marginTop: 55,
                width: '83%',
                justifyContent: 'center',
                alignItems: 'center',
                marginHorizontal: '8%',
            }} position="top" />
        </NavigationContainer>

    );
}

export default App;