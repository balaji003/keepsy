import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, SafeAreaView, KeyboardAvoidingView, Platform } from 'react-native';
import { StatusBar } from 'expo-status-bar';

export default function LoginScreen() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleLogin = () => {
        // TODO: Implement login logic
        console.log('Login attempt:', email);
    };

    return (
        <SafeAreaView className="flex-1 bg-gray-50">
            <StatusBar style="auto" />
            <KeyboardAvoidingView
                behavior={Platform.OS === "ios" ? "padding" : "height"}
                className="flex-1 justify-center px-6"
            >
                <View className="items-center mb-10">
                    <Text className="text-3xl font-bold text-gray-900 mb-2">Welcome Back</Text>
                    <Text className="text-gray-500 text-base">Sign in to your account</Text>
                </View>

                <View className="space-y-4">
                    <View>
                        <Text className="text-gray-700 font-medium mb-1 ml-1">Email</Text>
                        <TextInput
                            className="w-full bg-white border border-gray-300 rounded-lg px-4 py-3 text-gray-700 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200"
                            placeholder="you@example.com"
                            value={email}
                            onChangeText={setEmail}
                            autoCapitalize="none"
                            keyboardType="email-address"
                        />
                    </View>

                    <View>
                        <Text className="text-gray-700 font-medium mb-1 ml-1">Password</Text>
                        <TextInput
                            className="w-full bg-white border border-gray-300 rounded-lg px-4 py-3 text-gray-700 focus:border-indigo-500 focus:ring-2 focus:ring-indigo-200"
                            placeholder="••••••••"
                            value={password}
                            onChangeText={setPassword}
                            secureTextEntry
                        />
                    </View>

                    <TouchableOpacity className="items-end">
                        <Text className="text-indigo-600 font-medium">Forgot Password?</Text>
                    </TouchableOpacity>

                    <TouchableOpacity
                        className="w-full bg-indigo-600 rounded-lg py-3 items-center shadow-sm active:bg-indigo-700"
                        onPress={handleLogin}
                    >
                        <Text className="text-white font-bold text-lg">Sign In</Text>
                    </TouchableOpacity>
                </View>

                <View className="flex-row justify-center mt-8">
                    <Text className="text-gray-500">Don't have an account? </Text>
                    <TouchableOpacity>
                        <Text className="text-indigo-600 font-bold">Sign Up</Text>
                    </TouchableOpacity>
                </View>
            </KeyboardAvoidingView>
        </SafeAreaView>
    );
}
