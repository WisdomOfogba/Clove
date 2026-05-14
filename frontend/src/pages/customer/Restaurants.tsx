import React, { useEffect, useState } from 'react';
import { motion } from 'motion/react';
import { Link } from 'react-router-dom';
import { Icon } from '@iconify/react';
import { api } from '../../lib/api';
import { Button } from '../../components/ui/Button';
import { Badge } from '../../components/ui/Badge';
import { Input } from '../../components/ui/Input';

export function RestaurantsPage() {
  const [restaurants, setRestaurants] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [filter, setFilter] = useState('all');
  const [search, setSearch] = useState('');

  useEffect(() => {
    async function fetchRestaurants() {
      setLoading(true);
      const data = await api.getRestaurants();
      setRestaurants(data);
      setLoading(false);
    }
    fetchRestaurants();
  }, []);

  const filteredRestaurants = restaurants.filter(r => 
    r.name.toLowerCase().includes(search.toLowerCase()) && 
    (filter === 'all' || r.cuisine.toLowerCase() === filter.toLowerCase())
  );

  return (
    <div className="min-h-screen bg-muted/20 pb-20">
      <div className="bg-dark text-white pt-10 pb-24 rounded-b-[3rem] px-4">
        <div className="container mx-auto max-w-6xl space-y-8">
          <div className="flex flex-col md:flex-row justify-between items-center gap-6">
            <h1 className="text-4xl md:text-5xl font-black font-display tracking-tighter">
              Explore <span className="text-primary italic">Verified</span> Kitchens
            </h1>
            <div className="w-full md:w-96 text-dark">
              <Input 
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                placeholder="Search food or restaurant..." 
                icon={<Icon icon="solar:magnifer-bold-duotone" className="w-5 h-5 text-dark" />}
              />
            </div>
          </div>
          
          <div className="flex gap-4 overflow-x-auto no-scrollbar pb-2">
            {['All', 'Traditional', 'Fusion', 'BBQ', 'Fast Food'].map(c => (
              <button 
                key={c}
                onClick={() => setFilter(c.toLowerCase())}
                className={`px-6 py-2 rounded-full font-bold text-sm whitespace-nowrap transition-colors ${filter === c.toLowerCase() ? 'bg-primary text-dark' : 'bg-white/10 hover:bg-white/20'}`}
              >
                {c}
              </button>
            ))}
          </div>
        </div>
      </div>

      <div className="container mx-auto max-w-6xl px-4 -mt-10">
        {loading ? (
          <div className="flex justify-center py-20">
            <div className="w-12 h-12 border-4 border-primary border-t-transparent rounded-full animate-spin"></div>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            {filteredRestaurants.map(r => (
              <Link to={`/restaurants/${r.id}`} key={r.id}>
                <motion.div 
                  whileHover={{ y: -5 }}
                  className="bg-white rounded-[2.5rem] overflow-hidden shadow-[0_10px_40px_rgba(0,0,0,0.04)] hover:shadow-[0_10px_40px_rgba(0,0,0,0.08)] transition-shadow border border-muted"
                >
                  <div className="h-56 relative">
                    <img src={r.image} alt={r.name} className="w-full h-full object-cover" />
                    {r.isVerified && (
                      <div className="absolute top-4 left-4 bg-green-100 text-green-700 font-black text-[10px] uppercase tracking-widest px-3 py-1.5 rounded-full shadow-lg border border-green-200 flex items-center gap-1">
                        <Icon icon="solar:shield-check-bold" className="w-4 h-4" /> VERIFIED
                      </div>
                    )}
                  </div>
                  <div className="p-6 space-y-4">
                    <div className="flex justify-between items-start">
                      <div>
                        <h2 className="text-xl font-black text-dark">{r.name}</h2>
                        <p className="text-sm text-neutral font-medium">{r.cuisine} • {r.location}</p>
                      </div>
                      <div className="flex items-center gap-1 bg-slate-50 text-dark px-2 py-1 rounded-lg border border-slate-100">
                        <Icon icon="solar:star-bold" className="w-4 h-4 text-primary" />
                        <span className="font-bold text-sm">{r.rating}</span>
                      </div>
                    </div>
                    <div className="flex gap-2 flex-wrap">
                      {r.badges.map((b: string) => (
                        <Badge key={b} variant="success">{b}</Badge>
                      ))}
                    </div>
                    <div className="pt-4 border-t border-muted flex items-center justify-between">
                      <div className="flex flex-col">
                        <div className="flex items-center gap-2">
                          <Icon icon="solar:shield-check-bold" className="text-green-500 w-4 h-4" />
                          <span className="text-xs font-bold text-neutral uppercase tracking-widest">Clove Trust Score</span>
                        </div>
                        <span className="text-[9px] font-bold text-primary mt-0.5 uppercase tracking-wide">Powered by Clove AI</span>
                      </div>
                      <span className="font-black text-lg text-dark">{r.trustScore}%</span>
                    </div>
                  </div>
                </motion.div>
              </Link>
            ))}
          </div>
        )}
        
        {!loading && filteredRestaurants.length === 0 && (
          <div className="text-center py-20">
            <h3 className="text-2xl font-black text-dark">No kitchens found</h3>
            <p className="text-neutral mt-2">Try adjusting your filters or search term.</p>
          </div>
        )}
      </div>
    </div>
  );
}
