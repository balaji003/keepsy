import React from 'react';
import { View, Text, ScrollView, TouchableOpacity, SafeAreaView, StatusBar as RNStatusBar } from 'react-native';
import { StatusBar } from 'expo-status-bar';

// Mock Data
const stats = [
    { label: 'Total Products', value: '12' },
    { label: 'Total Value', value: '$4,250' },
    { label: 'Expiring Soon', value: '2' },
];

const recentProducts = [
    { id: 1, name: 'MacBook Pro', brand: 'Apple', purchaseDate: '2023-11-15', price: '$2400' },
    { id: 2, name: 'Sony XM5', brand: 'Sony', purchaseDate: '2024-01-20', price: '$350' },
    { id: 3, name: 'Keychron Q1', brand: 'Keychron', purchaseDate: '2024-02-10', price: '$180' },
];

import Card from '../components/Card';

export default function DashboardScreen({ navigation }: any) {
    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <View className="flex-1 px-4 pt-4" style={{ paddingTop: RNStatusBar.currentHeight }}>
                {/* Header */}
                <View className="flex-row justify-between items-center mb-6">
                    <View>
                        <Text className="text-subtext text-sm">Welcome back,</Text>
                        <Text className="text-white text-2xl font-bold">Balaji</Text>
                    </View>
                    <TouchableOpacity
                        className="bg-surface p-2 rounded-full border border-gray-800"
                        onPress={() => navigation.navigate('Profile')}
                    >
                        <View className="w-8 h-8 rounded-full bg-primary/20 items-center justify-center">
                            <Text className="text-primary font-bold">B</Text>
                        </View>
                    </TouchableOpacity>
                </View>

                <ScrollView showsVerticalScrollIndicator={false}>
                    {/* Stats Cards */}
                    <View className="flex-row justify-between mb-8">
                        <TouchableOpacity
                            className="w-[31%]"
                            onPress={() => navigation.navigate('ProductList')}
                        >
                            <Card className="items-center p-4">
                                <Text className="text-white text-xl font-bold mb-1">12</Text>
                                <Text className="text-subtext text-xs text-center">Products</Text>
                            </Card>
                        </TouchableOpacity>

                        <View className="w-[31%]">
                            <Card className="items-center p-4">
                                <Text className="text-white text-xl font-bold mb-1">$4.2k</Text>
                                <Text className="text-subtext text-xs text-center">Value</Text>
                            </Card>
                        </View>

                        <TouchableOpacity
                            className="w-[31%]"
                            onPress={() => navigation.navigate('ProductList')}
                        >
                            <Card className="items-center p-4 border-red-900/50 bg-red-900/10">
                                <Text className="text-red-400 text-xl font-bold mb-1">2</Text>
                                <Text className="text-red-400/70 text-xs text-center">Expiring</Text>
                            </Card>
                        </TouchableOpacity>
                    </View>

                    {/* Quick Actions */}
                    <View className="mb-8">
                        <Text className="text-white text-lg font-bold mb-4">Quick Actions</Text>
                        <View className="flex-row space-x-4">
                            <TouchableOpacity
                                className="flex-1 bg-primary p-4 rounded-xl items-center flex-row justify-center space-x-2"
                                onPress={() => navigation.navigate('AddProduct')}
                            >
                                <Text className="text-white font-bold text-base">+ Add Product</Text>
                            </TouchableOpacity>
                            <TouchableOpacity
                                className="flex-1 bg-surface border border-gray-800 p-4 rounded-xl items-center flex-row justify-center space-x-2"
                                onPress={() => navigation.navigate('BillList')}
                            >
                                <Text className="text-white font-bold text-base">Scan Bill</Text>
                            </TouchableOpacity>
                        </View>
                    </View>

                    {/* Recent Products */}
                    <View className="mb-8">
                        <View className="flex-row justify-between items-center mb-4">
                            <Text className="text-white text-lg font-bold">Recent Products</Text>
                            <TouchableOpacity onPress={() => navigation.navigate('ProductList')}>
                                <Text className="text-primary text-sm">View All</Text>
                            </TouchableOpacity>
                        </View>
                        <View className="space-y-3">
                            {recentProducts.map((product) => (
                                <TouchableOpacity key={product.id} className="bg-surface p-4 rounded-xl border border-gray-800 flex-row justify-between items-center">
                                    <View>
                                        <Text className="text-white font-bold text-base">{product.name}</Text>
                                        <Text className="text-subtext text-sm">{product.brand}</Text>
                                    </View>
                                    <View className="items-end">
                                        <Text className="text-primary font-bold text-base">{product.price}</Text>
                                        <Text className="text-subtext text-xs">{product.purchaseDate}</Text>
                                    </View>
                                </TouchableOpacity>
                            ))}
                        </View>
                    </View>
                </ScrollView>
            </View>
        </SafeAreaView>
    );
}
