import React, { useState } from 'react';
import { View, Text, FlatList, TouchableOpacity, SafeAreaView, StatusBar as RNStatusBar } from 'react-native';
import { StatusBar } from 'expo-status-bar';

// Mock Data
const mockProducts = [
    { id: 1, name: 'MacBook Pro', brand: 'Apple', purchaseDate: '2023-11-15', price: '$2400', warrantyEnd: '2024-11-15' },
    { id: 2, name: 'Sony XM5', brand: 'Sony', purchaseDate: '2024-01-20', price: '$350', warrantyEnd: '2025-01-20' },
    { id: 3, name: 'Keychron Q1', brand: 'Keychron', purchaseDate: '2024-02-10', price: '$180', warrantyEnd: '2025-02-10' },
    { id: 4, name: 'LG Monitor', brand: 'LG', purchaseDate: '2023-05-01', price: '$400', warrantyEnd: '2026-05-01' },
];

import Card from '../components/Card';

export default function ProductListScreen({ navigation }: any) {
    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <View className="flex-1 px-4 mb-4" style={{ paddingTop: RNStatusBar.currentHeight }}>
                <View className="flex-row justify-between items-center mb-6 pt-4">
                    <TouchableOpacity onPress={() => navigation.goBack()} className="p-2">
                        <Text className="text-white text-lg">Back</Text>
                    </TouchableOpacity>
                    <Text className="text-white text-xl font-bold">My Products</Text>
                    <TouchableOpacity onPress={() => navigation.navigate('AddProduct')} className="p-2">
                        <Text className="text-primary text-lg">Add</Text>
                    </TouchableOpacity>
                </View>

                <FlatList
                    data={mockProducts}
                    keyExtractor={(item) => item.id.toString()}
                    contentContainerStyle={{ paddingBottom: 20 }}
                    showsVerticalScrollIndicator={false}
                    renderItem={({ item }) => (
                        <TouchableOpacity onPress={() => navigation.navigate('ProductDetails', { product: item })}>
                            <Card className="mb-3">
                                <View className="flex-row justify-between items-start mb-2">
                                    <View>
                                        <Text className="text-white font-bold text-lg">{item.name}</Text>
                                        <Text className="text-subtext text-sm">{item.brand}</Text>
                                    </View>
                                    <Text className="text-primary font-bold text-lg">{item.price}</Text>
                                </View>
                                <View className="h-[1px] bg-gray-800 my-2" />
                                <View className="flex-row justify-between">
                                    <Text className="text-subtext text-xs">Bought: {item.purchaseDate}</Text>
                                    <Text className="text-subtext text-xs text-secondary">Warranty: {item.warrantyEnd}</Text>
                                </View>
                            </Card>
                        </TouchableOpacity>
                    )}
                />
            </View>
        </SafeAreaView>
    );
}
