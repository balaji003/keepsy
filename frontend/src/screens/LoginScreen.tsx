import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, SafeAreaView, KeyboardAvoidingView, Platform } from 'react-native';
import { StatusBar } from 'expo-status-bar';

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

                <View className="space-y-4">
                    <View>
                        <Text className="text-gray-400 font-medium mb-1 ml-1">Email</Text>
                        <TextInput
                            className="w-full bg-surface border border-gray-800 rounded-lg px-4 py-3 text-white focus:border-primary focus:ring-1 focus:ring-primary placeholder:text-gray-600"
                            placeholder="you@example.com"
                            placeholderTextColor="#6b7280"
                            value={email}
                            onChangeText={setEmail}
                            autoCapitalize="none"
                            keyboardType="email-address"
                        />
                    </View>

                    <View>
                        <Text className="text-gray-400 font-medium mb-1 ml-1">Password</Text>
                        <TextInput
                            className="w-full bg-surface border border-gray-800 rounded-lg px-4 py-3 text-white focus:border-primary focus:ring-1 focus:ring-primary placeholder:text-gray-600"
                            placeholder="••••••••"
                            placeholderTextColor="#6b7280"
                            value={password}
                            onChangeText={setPassword}
                            secureTextEntry
                        />
                    </View>

                    <TouchableOpacity className="items-end">
                        <Text className="text-primary font-medium">Forgot Password?</Text>
                    </TouchableOpacity>

                    <TouchableOpacity
                        className="w-full bg-primary rounded-lg py-3 items-center shadow-lg shadow-primary/20 active:bg-secondary"
                        onPress={handleLogin}
                    >
                        <Text className="text-white font-bold text-lg">Sign In</Text>
                    </TouchableOpacity>
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
