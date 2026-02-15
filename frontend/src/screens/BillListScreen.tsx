import React, { useState } from 'react';
import { View, Text, FlatList, TouchableOpacity, SafeAreaView, StatusBar as RNStatusBar } from 'react-native';
import { StatusBar } from 'expo-status-bar';

// Mock Data
const mockBills = [
    { id: 1, name: 'MacBook Pro Receipt', date: '2023-11-15', type: 'PDF' },
    { id: 2, name: 'Sony XM5 Invoice', date: '2024-01-20', type: 'Image' },
    { id: 3, name: 'Uber Ride', date: '2024-02-14', type: 'PDF' },
];

import Card from '../components/Card';

export default function BillListScreen({ navigation }: any) {
    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <View className="flex-1 px-4 mb-4" style={{ paddingTop: RNStatusBar.currentHeight }}>
                <View className="flex-row justify-between items-center mb-6 pt-4">
                    <TouchableOpacity onPress={() => navigation.goBack()} className="p-2">
                        <Text className="text-white text-lg">Back</Text>
                    </TouchableOpacity>
                    <Text className="text-white text-xl font-bold">My Bills</Text>
                    <TouchableOpacity onPress={() => navigation.navigate('AddProduct')} className="p-2">
                        <Text className="text-primary text-lg">Add</Text>
                    </TouchableOpacity>
                </View>

                <FlatList
                    data={mockBills}
                    keyExtractor={(item) => item.id.toString()}
                    contentContainerStyle={{ paddingBottom: 20 }}
                    showsVerticalScrollIndicator={false}
                    renderItem={({ item }) => (
                        <TouchableOpacity onPress={() => navigation.navigate('BillDetails', { bill: item })}>
                            <Card className="mb-3 flex-row items-center justify-between p-4">
                                <View className="flex-row items-center space-x-4">
                                    <View className="w-10 h-10 bg-gray-800 rounded-lg items-center justify-center">
                                        <Text className="text-gray-400 font-bold">{item.type}</Text>
                                    </View>
                                    <View>
                                        <Text className="text-white font-bold text-base">{item.name}</Text>
                                        <Text className="text-subtext text-xs">{item.date}</Text>
                                    </View>
                                </View>
                                <Text className="text-primary text-sm">View</Text>
                            </Card>
                        </TouchableOpacity>
                    )}
                />
            </View>
        </SafeAreaView>
    );
}
