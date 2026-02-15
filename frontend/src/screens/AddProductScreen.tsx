import React, { useState } from 'react';
import { View, Text, ScrollView, Platform, KeyboardAvoidingView, TouchableOpacity, SafeAreaView } from 'react-native';
import { StatusBar } from 'expo-status-bar';
import Button from '../components/Button';
import Input from '../components/Input';
import Card from '../components/Card';

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
                        <Input
                            label="Product Name"
                            placeholder="e.g. MacBook Pro"
                            value={name}
                            onChangeText={setName}
                        />

                        <Input
                            label="Brand"
                            placeholder="e.g. Apple"
                            value={brand}
                            onChangeText={setBrand}
                        />

                        <Input
                            label="Price"
                            placeholder="0.00"
                            value={price}
                            onChangeText={setPrice}
                            keyboardType="numeric"
                        />

                        <Input
                            label="Purchase Date"
                            placeholder="YYYY-MM-DD"
                            value={date}
                            onChangeText={setDate}
                        />

                        <TouchableOpacity
                            className="mt-4 bg-surface border border-gray-800 border-dashed rounded-xl p-8 items-center justify-center space-y-2"
                        >
                            <View className="w-12 h-12 rounded-full bg-gray-800 items-center justify-center">
                                <Text className="text-primary text-2xl">+</Text>
                            </View>
                            <Text className="text-gray-400">Upload Receipt / Bill</Text>
                        </TouchableOpacity>

                        <Button
                            title="Save Product"
                            onPress={() => console.log('Save')}
                            className="mt-4"
                        />
                    </View>
                </ScrollView>
            </KeyboardAvoidingView>
        </SafeAreaView>
    );
}
