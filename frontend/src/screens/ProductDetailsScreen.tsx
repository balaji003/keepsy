import React from 'react';
import { View, Text, ScrollView, TouchableOpacity, SafeAreaView, StatusBar as RNStatusBar } from 'react-native';
import { StatusBar } from 'expo-status-bar';
import Card from '../components/Card';
import Button from '../components/Button';

export default function ProductDetailsScreen({ route, navigation }: any) {
    const { product } = route.params || {};

    // Mock data if no product is passed (for development/testing)
    const item = product || {
        id: 1,
        name: 'MacBook Pro',
        brand: 'Apple',
        purchaseDate: '2023-11-15',
        price: '$2400',
        warrantyEnd: '2024-11-15',
        hasBill: true,
    };

    return (
        <SafeAreaView className="flex-1 bg-background">
            <StatusBar style="light" />
            <View className="flex-1 px-4 mb-4" style={{ paddingTop: RNStatusBar.currentHeight }}>
                <View className="flex-row justify-between items-center mb-6 pt-4">
                    <TouchableOpacity onPress={() => navigation.goBack()} className="p-2">
                        <Text className="text-white text-lg">Back</Text>
                    </TouchableOpacity>
                    <Text className="text-white text-lg font-bold" numberOfLines={1}>Product Details</Text>
                    <TouchableOpacity onPress={() => console.log('Edit')} className="p-2">
                        <Text className="text-primary text-lg">Edit</Text>
                    </TouchableOpacity>
                </View>

                <ScrollView showsVerticalScrollIndicator={false}>
                    <View className="items-center mb-8">
                        <View className="w-24 h-24 bg-surface rounded-full items-center justify-center border border-gray-800 mb-4">
                            <Text className="text-primary text-4xl font-bold">{item.name.charAt(0)}</Text>
                        </View>
                        <Text className="text-white text-2xl font-bold text-center">{item.name}</Text>
                        <Text className="text-subtext text-lg">{item.brand}</Text>
                    </View>

                    <Card className="mb-6">
                        <Text className="text-white font-bold text-lg mb-4">Information</Text>

                        <View className="space-y-4">
                            <View className="flex-row justify-between">
                                <Text className="text-subtext">Price</Text>
                                <Text className="text-white font-medium">{item.price}</Text>
                            </View>
                            <View className="h-[1px] bg-gray-800" />

                            <View className="flex-row justify-between">
                                <Text className="text-subtext">Purchase Date</Text>
                                <Text className="text-white font-medium">{item.purchaseDate}</Text>
                            </View>
                            <View className="h-[1px] bg-gray-800" />

                            <View className="flex-row justify-between">
                                <Text className="text-subtext">Warranty Ends</Text>
                                <Text className="text-white font-medium">{item.warrantyEnd || 'N/A'}</Text>
                            </View>
                        </View>
                    </Card>

                    <Card className="mb-8">
                        <View className="flex-row justify-between items-center">
                            <View>
                                <Text className="text-white font-bold text-lg">Bill / Receipt</Text>
                                <Text className="text-subtext text-xs mt-1">
                                    {item.hasBill ? 'Receipt linked' : 'No receipt uploaded'}
                                </Text>
                            </View>
                            {item.hasBill && (
                                <TouchableOpacity
                                    className="bg-gray-800 px-4 py-2 rounded-lg"
                                    onPress={() => navigation.navigate('BillDetails', { billId: item.id })}
                                >
                                    <Text className="text-white text-sm">View</Text>
                                </TouchableOpacity>
                            )}
                        </View>
                    </Card>

                    <Button
                        title="Delete Product"
                        variant="outline"
                        className="border-red-500 mb-8"
                        onPress={() => console.log('Delete')}
                    />
                    {/* Text style override for delete button */}
                    <Text className="text-red-500 font-bold text-lg text-center absolute bottom-12 left-0 right-0 pointer-events-none">Delete Product</Text>
                </ScrollView>
            </View>
        </SafeAreaView>
    );
}
