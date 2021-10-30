// In App.js in a new project

import * as React from 'react';
import { View, Text } from 'react-native';
import {NavigationContainer} from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import SignUp from './Pages/SignUp';
import SignUpForm from './Pages/SignUpForm';
import CheckEmailToken from './Pages/CheckEmailToken';
import SuccessSignUp from './Pages/SuccessSignUp';

function HomeScreen() {
    return (
        <View style={{ flex: 1, alignItems: 'center', justifyContent: 'center' }}>
            <Text>Home Screen</Text>
        </View>
    );
}

const Stack = createNativeStackNavigator();

function App() {
    return (
        <NavigationContainer>
            <Stack.Navigator screenOptions={{headerShown: false}}>
                <Stack.Screen name="SignUp" component={SignUp} />
                <Stack.Screen name="SignUpForm" component={SignUpForm} />
                <Stack.Screen name="CheckEmailToken" component={CheckEmailToken} />
                <Stack.Screen name="SuccessSignUp" component={SuccessSignUp} />
            </Stack.Navigator>
        </NavigationContainer>
    );
}

export default App;