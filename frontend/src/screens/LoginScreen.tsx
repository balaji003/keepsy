import React, { useState } from 'react';
import { View, Text, SafeAreaView, KeyboardAvoidingView, Platform, TouchableOpacity } from 'react-native';
import { StatusBar } from 'expo-status-bar';
import Button from '../components/Button';
import Input from '../components/Input';

export default function LoginScreen({ navigation }: any) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = () => {
        // TODO: Implement login logic
        console.log('Login attempt:', email);
        navigation.replace('Dashboard');
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
                        title="Sign In"
                        onPress={handleLogin}
                        variant="primary"
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
