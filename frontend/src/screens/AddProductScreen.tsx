import React, { useState } from 'react';
import { View, Text, TextInput, TouchableOpacity, SafeAreaView, ScrollView, Platform, KeyboardAvoidingView } from 'react-native';
import { StatusBar } from 'expo-status-bar';

export default function AddProductScreen({ navigation }: any) {
    const [name, setName] = useState('');
    const [brand, setBrand] = useState('');
    const [price, setPrice] = useState('');
    const [date, setDate] = useState('');

    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <KeyboardAvoidingView
                behavior={Platform.OS === "ios" ? "padding" : "height"}
                className="flex-1"
            >
                <ScrollView className="flex-1 px-4 pt-4">
                    <View className="flex-row justify-between items-center mb-8">
                        <TouchableOpacity onPress={() => navigation.goBack()} className="p-2">
                            <Text className="text-gray-400 text-lg">Cancel</Text>
                        </TouchableOpacity>
                        <Text className="text-white text-xl font-bold">Add Product</Text>
                        <TouchableOpacity onPress={() => console.log('Save')} className="p-2">
                            <Text className="text-primary text-lg font-bold">Save</Text>
                        </TouchableOpacity>
                    </View>

                    <View className="space-y-6">
                        <View>
                            <Text className="text-gray-400 font-medium mb-2 ml-1">Product Name</Text>
                            <TextInput
                                className="w-full bg-surface border border-gray-800 rounded-lg px-4 py-3 text-white focus:border-primary focus:ring-1 focus:ring-primary placeholder:text-gray-600"
                                placeholder="e.g. MacBook Pro"
                                placeholderTextColor="#6b7280"
                                value={name}
                                onChangeText={setName}
                            />
                        </View>

                        <View>
                            <Text className="text-gray-400 font-medium mb-2 ml-1">Brand</Text>
                            <TextInput
                                className="w-full bg-surface border border-gray-800 rounded-lg px-4 py-3 text-white focus:border-primary focus:ring-1 focus:ring-primary placeholder:text-gray-600"
                                placeholder="e.g. Apple"
                                placeholderTextColor="#6b7280"
                                value={brand}
                                onChangeText={setBrand}
                            />
                        </View>

                        <View>
                            <Text className="text-gray-400 font-medium mb-2 ml-1">Price</Text>
                            <TextInput
                                className="w-full bg-surface border border-gray-800 rounded-lg px-4 py-3 text-white focus:border-primary focus:ring-1 focus:ring-primary placeholder:text-gray-600"
                                placeholder="0.00"
                                placeholderTextColor="#6b7280"
                                value={price}
                                onChangeText={setPrice}
                                keyboardType="numeric"
                            />
                        </View>

                        <View>
                            <Text className="text-gray-400 font-medium mb-2 ml-1">Purchase Date</Text>
                            <TextInput
                                className="w-full bg-surface border border-gray-800 rounded-lg px-4 py-3 text-white focus:border-primary focus:ring-1 focus:ring-primary placeholder:text-gray-600"
                                placeholder="YYYY-MM-DD"
                                placeholderTextColor="#6b7280"
                                value={date}
                                onChangeText={setDate}
                            />
                        </View>

                        <TouchableOpacity
                            className="mt-4 bg-surface border border-gray-800 border-dashed rounded-xl p-8 items-center justify-center space-y-2"
                        >
                            <View className="w-12 h-12 rounded-full bg-gray-800 items-center justify-center">
                                <Text className="text-primary text-2xl">+</Text>
                            </View>
                            <Text className="text-gray-400">Upload Receipt / Bill</Text>
                        </TouchableOpacity>
                    </View>
                </ScrollView>
            </KeyboardAvoidingView>
        </SafeAreaView>
    );
}
