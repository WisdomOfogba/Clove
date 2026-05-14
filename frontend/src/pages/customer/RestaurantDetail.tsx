import React, { useEffect, useState } from 'react';
import { useParams, Link } from 'react-router-dom';
import { motion } from 'motion/react';
import { Icon } from '@iconify/react';
import { api } from '../../lib/api';
import { Button } from '../../components/ui/Button';
import { Badge } from '../../components/ui/Badge';
import { useCart } from '../../contexts/CartContext';

export function RestaurantDetailPage() {
  const { id } = useParams();
  const [restaurant, setRestaurant] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const { addItem } = useCart();

  useEffect(() => {
    async function loadRestaurant() {
      if (!id) return;
      const data = await api.getRestaurant(id);
      setRestaurant(data);
      setLoading(false);
    }
    loadRestaurant();
  }, [id]);

  if (loading) {
    return <div className="flex justify-center items-center h-[60vh]"><div className="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin"></div></div>;
  }

  if (!restaurant) {
    return <div className="text-center py-20"><h1 className="text-4xl font-black text-dark">Kitchen not found.</h1></div>;
  }

  return (
    <div className="min-h-screen bg-muted/20 pb-20">
      {/* Header Banner */}
      <div className="h-[40vh] min-h-[300px] relative">
        <img src={restaurant.image} alt={restaurant.name} className="w-full h-full object-cover" />
        <div className="absolute inset-0 bg-gradient-to-t from-dark via-dark/40 to-transparent" />
        <div className="absolute top-4 left-4">
          <Link to="/restaurants" className="w-10 h-10 bg-white/20 backdrop-blur-md rounded-full flex items-center justify-center text-white hover:bg-white/40 transition-colors">
            <Icon icon="solar:arrow-left-bold" className="w-5 h-5" />
          </Link>
        </div>
      </div>

      <div className="container mx-auto px-4 max-w-5xl -mt-20 relative z-10 space-y-10">
        
        {/* Info Card */}
        <div className="bg-white rounded-[3rem] p-8 sm:p-12 shadow-2xl border border-muted flex flex-col md:flex-row gap-8 justify-between items-start">
          <div className="space-y-4 max-w-2xl">
            <div className="flex items-center gap-3">
              <Badge variant="success" className="px-4 py-1.5"><Icon icon="solar:shield-check-bold" className="mr-1"/> Clove Verified</Badge>
              <span className="text-sm font-bold text-neutral">{restaurant.cuisine}</span>
            </div>
            <h1 className="text-5xl font-black text-dark font-display">{restaurant.name}</h1>
            <p className="text-lg font-medium text-neutral flex items-center gap-2">
              <Icon icon="solar:map-point-bold-duotone" className="text-primary w-5 h-5"/>
              {restaurant.location}
            </p>
          </div>

          <div className="bg-slate-50 p-6 rounded-[2rem] border border-muted text-center min-w-[200px]">
             <p className="text-xs font-black uppercase tracking-widest text-neutral mb-2">Clove Trust Score</p>
             <div className="flex justify-center items-end gap-1 mb-1">
               <span className="text-5xl font-black text-primary">{restaurant.trustScore}</span>
               <span className="text-xl font-bold text-neutral mb-1">%</span>
             </div>
             <p className="text-[10px] font-bold text-primary tracking-wide uppercase mt-1">Powered by Clove AI</p>
             <p className="text-xs font-medium text-slate-500 mt-2">Verified via NIN & Smile ID</p>
          </div>
        </div>

        {/* Menu Section */}
        <div className="space-y-6">
          <h2 className="text-3xl font-black text-dark">Menu Items</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {restaurant.menu?.map((item: any) => (
              <div key={item.id} className="bg-white p-4 rounded-[2rem] flex gap-4 shadow-sm hover:shadow-md transition-shadow border border-slate-100">
                <div className="w-32 h-32 rounded-2xl overflow-hidden shrink-0">
                  <img src={item.image} alt={item.name} className="w-full h-full object-cover" />
                </div>
                <div className="flex-1 flex flex-col justify-between py-2">
                  <div>
                    <h3 className="text-lg font-black text-dark">{item.name}</h3>
                    <p className="text-sm font-medium text-neutral mt-1 line-clamp-2">{item.description}</p>
                  </div>
                  <div className="flex justify-between items-center mt-4">
                    <span className="font-black text-dark text-lg">₦{item.price}</span>
                    <Button size="sm" onClick={() => addItem({ id: item.id, name: item.name, price: item.price, vendorName: restaurant.name, quantity: 1 })}>
                      Add to Cart
                    </Button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>

      </div>
    </div>
  );
}
