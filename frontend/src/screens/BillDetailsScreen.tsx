import React from 'react';
import { View, Text, ScrollView, TouchableOpacity, SafeAreaView, Image, StatusBar as RNStatusBar } from 'react-native';
import { StatusBar } from 'expo-status-bar';
import Card from '../components/Card';
import Button from '../components/Button';

export default function BillDetailsScreen({ route, navigation }: any) {
    const { bill } = route.params || {};

    const item = bill || {
        id: 1,
        name: 'MacBook Pro Receipt',
        date: '2023-11-15',
        amount: '$2400',
        category: 'Electronics',
        image: 'https://via.placeholder.com/400x600', // Mock image
    };

    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <View className="flex-1 px-4 mb-4" style={{ paddingTop: RNStatusBar.currentHeight }}>
                <View className="flex-row justify-between items-center mb-6 pt-4">
                    <TouchableOpacity onPress={() => navigation.goBack()} className="p-2">
                        <Text className="text-white text-lg">Back</Text>
                    </TouchableOpacity>
                    <Text className="text-white text-lg font-bold">Bill Details</Text>
                    <TouchableOpacity onPress={() => console.log('Share')} className="p-2">
                        <Text className="text-primary text-lg">Share</Text>
                    </TouchableOpacity>
                </View>

                <ScrollView showsVerticalScrollIndicator={false}>
                    <Card className="mb-6 p-2">
                        <View className="h-96 w-full bg-gray-800 rounded-xl overflow-hidden items-center justify-center">
                            {/* Mock Image Viewer */}
                            <Image
                                source={{ uri: item.image }}
                                className="w-full h-full"
                                resizeMode="contain"
                            />
                        </View>
                    </Card>

                    <Card className="mb-6">
                        <Text className="text-white font-bold text-lg mb-4">Details</Text>

                        <View className="space-y-4">
                            <View className="flex-row justify-between">
                                <Text className="text-subtext">Item</Text>
                                <Text className="text-white font-medium">{item.name}</Text>
                            </View>
                            <View className="h-[1px] bg-gray-800" />

                            <View className="flex-row justify-between">
                                <Text className="text-subtext">Category</Text>
                                <Text className="text-white font-medium">{item.category}</Text>
                            </View>
                            <View className="h-[1px] bg-gray-800" />

                            <View className="flex-row justify-between">
                                <Text className="text-subtext">Date</Text>
                                <Text className="text-white font-medium">{item.date}</Text>
                            </View>
                            <View className="h-[1px] bg-gray-800" />

                            <View className="flex-row justify-between">
                                <Text className="text-subtext">Amount</Text>
                                <Text className="text-white font-medium">{item.amount}</Text>
                            </View>
                        </View>
                    </Card>

                    <Button
                        title="Delete Bill"
                        variant="outline"
                        className="border-red-500 mb-8"
                        onPress={() => console.log('Delete Bill')}
                    />
                </ScrollView>
            </View>
        </SafeAreaView>
    );
}
