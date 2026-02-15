import React, { useState } from 'react';
import { View, Text, SafeAreaView, KeyboardAvoidingView, Platform, TouchableOpacity, Alert } from 'react-native';
import { StatusBar } from 'expo-status-bar';
import Button from '../components/Button';
import Input from '../components/Input';

export default function SignUpScreen({ navigation }: any) {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleSignUp = async () => {
        if (!name || !email || !password) {
            Alert.alert('Error', 'Please fill in all fields');
            return;
        }

        setIsLoading(true);
        try {
            // Use local IP if testing on device, localhost for simulator/web
            const response = await fetch('http://localhost:8080/auth/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    name: name,
                    email: email,
                    password: password,
                }),
            });

            const data = await response.json();

            if (response.ok) {
                console.log('Signup successful:', data);
                Alert.alert('Success', 'Account created! Please sign in.', [
                    { text: 'OK', onPress: () => navigation.navigate('Login') }
                ]);
            } else {
                Alert.alert('Signup Failed', data.message || 'Could not create account');
            }
        } catch (error) {
            console.error('Signup error:', error);
            Alert.alert('Error', 'Something went wrong. Please try again.');
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <KeyboardAvoidingView
                behavior={Platform.OS === "ios" ? "padding" : "height"}
                className="flex-1 justify-center px-6"
            >
                <View className="items-center mb-10">
                    <Text className="text-3xl font-bold text-white mb-2">Create Account</Text>
                    <Text className="text-gray-400 text-base">Sign up to get started</Text>
                </View>

                <View className="space-y-6">
                    <Input
                        label="Full Name"
                        placeholder="John Doe"
                        value={name}
                        onChangeText={setName}
                    />

                    <Input
                        label="Email"
                        placeholder="you@example.com"
                        value={email}
                        onChangeText={setEmail}
                        autoCapitalize="none"
                        keyboardType="email-address"
                    />

                    <Input
                        label="Password"
                        placeholder="••••••••"
                        value={password}
                        onChangeText={setPassword}
                        secureTextEntry
                    />

                    <Button
                        title={isLoading ? "Creating Account..." : "Sign Up"}
                        onPress={handleSignUp}
                        variant="primary"
                        disabled={isLoading}
                    />
                </View>

                <View className="flex-row justify-center mt-8">
                    <Text className="text-gray-500">Already have an account? </Text>
                    <TouchableOpacity onPress={() => navigation.navigate('Login')}>
                        <Text className="text-primary font-bold">Sign In</Text>
                    </TouchableOpacity>
                </View>
            </KeyboardAvoidingView>
        </SafeAreaView>
    );
}
