import React, { useState } from 'react';
import { Icon } from '@iconify/react';
import { VendorLayout } from '../../components/layout/VendorLayout';
import { Button } from '../../components/ui/Button';

export function VendorMeals() {
  const [meals] = useState([
    { id: 1, name: 'Party Jollof Rice', price: 3500, category: 'Rice', status: 'active', img: 'https://images.unsplash.com/photo-1565299507177-b0ac66763828?auto=format&fit=crop&q=80' },
    { id: 2, name: 'Egusi Soup & Pounded Yam', price: 4500, category: 'Swallow', status: 'active', img: 'https://images.unsplash.com/photo-1544025162-d76694265947?auto=format&fit=crop&q=80' },
    { id: 3, name: 'Asun (Spicy Goat Meat)', price: 5000, category: 'Protein', status: 'draft', img: 'https://images.unsplash.com/photo-1504674900247-0877df9cc836?auto=format&fit=crop&q=80' },
  ]);

  return (
    <VendorLayout>
       <div className="space-y-8">
          <div className="flex justify-between items-center bg-white p-6 rounded-[2.5rem] border border-slate-200 shadow-sm">
             <div>
                <h2 className="text-2xl font-black text-dark">Menu Management</h2>
                <p className="text-sm font-medium text-neutral">Add or edit your meals.</p>
             </div>
             <Button className="rounded-full px-6 flex items-center gap-2 h-12">
                <Icon icon="solar:add-circle-bold" className="w-5 h-5" />
                Add Meal
             </Button>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
             {meals.map(meal => (
                <div key={meal.id} className="bg-white p-4 rounded-[2.5rem] border border-slate-200 shadow-sm flex gap-4 group">
                   <div className="w-32 h-32 rounded-[2rem] overflow-hidden bg-slate-100 shrink-0">
                      <img src={meal.img} className="w-full h-full object-cover" alt={meal.name} />
                   </div>
                   <div className="flex-1 flex flex-col justify-center">
                      <div className="flex justify-between items-start mb-1">
                         <span className="text-[10px] font-black uppercase tracking-widest text-primary bg-primary/10 px-2 py-1 rounded-md">{meal.category}</span>
                         <span className={`text-[10px] font-black uppercase tracking-widest px-2 py-1 rounded-md ${meal.status === 'active' ? 'text-green-600 bg-green-50' : 'text-neutral bg-slate-100'}`}>
                           {meal.status}
                         </span>
                      </div>
                      <h3 className="font-black text-dark text-lg leading-tight mb-2">{meal.name}</h3>
                      <p className="font-bold text-neutral">₦{meal.price.toLocaleString()}</p>
                      
                      <div className="flex gap-2 mt-4 opacity-0 group-hover:opacity-100 transition-opacity">
                         <button className="flex-1 bg-slate-100 hover:bg-slate-200 text-dark text-xs font-bold py-2 rounded-xl transition-colors">Edit</button>
                         <button className="flex-1 bg-red-50 hover:bg-red-100 text-red-600 text-xs font-bold py-2 rounded-xl transition-colors">Delete</button>
                      </div>
                   </div>
                </div>
             ))}
          </div>
       </div>
    </VendorLayout>
  );
}
