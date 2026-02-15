import React, { useState } from 'react';
import { View, Text, SafeAreaView, KeyboardAvoidingView, Platform, TouchableOpacity, Alert } from 'react-native';
import { StatusBar } from 'expo-status-bar';
import Button from '../components/Button';
import Input from '../components/Input';

export default function LoginScreen({ navigation }: any) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [isLoading, setIsLoading] = useState(false);

    const handleLogin = async () => {
        if (!email || !password) {
            Alert.alert('Error', 'Please enter both email and password');
            return;
        }

        setIsLoading(true);
        try {
            // Use local IP if testing on device, localhost for simulator/web
            const response = await fetch('http://localhost:8080/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    identifier: email,
                    password: password,
                }),
            });

            const data = await response.json();

            if (response.ok) {
                console.log('Login successful:', data);
                navigation.replace('Dashboard');
            } else {
                Alert.alert('Login Failed', data.message || 'Invalid credentials');
            }
        } catch (error) {
            console.error('Login error:', error);
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
                    <Text className="text-3xl font-bold text-white mb-2">Welcome Back</Text>
                    <Text className="text-gray-400 text-base">Sign in to your account</Text>
                </View>

                <View className="space-y-6">
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

                    <TouchableOpacity className="items-end">
                        <Text className="text-primary font-medium">Forgot Password?</Text>
                    </TouchableOpacity>

                    <Button
                        title={isLoading ? "Signing In..." : "Sign In"}
                        onPress={handleLogin}
                        variant="primary"
                        disabled={isLoading}
                    />
                </View>

                <View className="flex-row justify-center mt-8">
                    <Text className="text-gray-500">Don't have an account? </Text>
                    <TouchableOpacity>
                        <Text className="text-primary font-bold">Sign Up</Text>
                    </TouchableOpacity>
                </View>
            </KeyboardAvoidingView>
        </SafeAreaView>
    );
}
