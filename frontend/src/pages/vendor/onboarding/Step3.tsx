import { useOnboardingStore } from '../../../store';
import { Icon } from '@iconify/react';
import { useState, useRef, useEffect } from 'react';
import { motion, AnimatePresence } from 'motion/react';
import { cn } from '../../../lib/utils';

export function Step3() {
  const { data, updateProof, setStep } = useOnboardingStore();
  const [activeCamera, setActiveCamera] = useState<'kitchen' | 'meal' | 'exterior' | 'selfie' | null>(null);
  const [capturedPhotos, setCapturedPhotos] = useState<Record<string, string[]>>({
    kitchen: data.proof.kitchenPhotos,
    meal: data.proof.mealPhotos,
    exterior: data.proof.exteriorPhoto ? [data.proof.exteriorPhoto] : [],
    selfie: data.proof.selfie ? [data.proof.selfie] : []
  });

  const [geoStatus, setGeoStatus] = useState<'idle' | 'loading' | 'success' | 'error'>('idle');
  const videoRef = useRef<HTMLVideoElement>(null);
  const [countdown, setCountdown] = useState<number | null>(null);

  useEffect(() => {
    if (activeCamera && videoRef.current) {
      navigator.mediaDevices.getUserMedia({ video: { facingMode: activeCamera === 'selfie' ? 'user' : 'environment' } })
        .then(stream => {
          if (videoRef.current) videoRef.current.srcObject = stream;
        });
    }
    return () => {
      if (videoRef.current?.srcObject) {
        (videoRef.current.srcObject as MediaStream).getTracks().forEach(t => t.stop());
      }
    };
  }, [activeCamera]);

  const capturePhoto = () => {
    if (!videoRef.current || !activeCamera) return;
    
    if (activeCamera === 'selfie' && countdown === null) {
      setCountdown(3);
      const timer = setInterval(() => {
        setCountdown(prev => {
          if (prev === 1) {
            clearInterval(timer);
            const canvas = document.createElement('canvas');
            canvas.width = videoRef.current!.videoWidth;
            canvas.height = videoRef.current!.videoHeight;
            canvas.getContext('2d')?.drawImage(videoRef.current!, 0, 0);
            const dataUrl = canvas.toDataURL('image/jpeg');
            updatePhotoState(activeCamera, dataUrl);
            setCountdown(null);
            setTimeout(() => setActiveCamera(null), 800);
            return null;
          }
          return prev ? prev - 1 : null;
        });
      }, 1000);
      return;
    }

    const canvas = document.createElement('canvas');
    canvas.width = videoRef.current.videoWidth;
    canvas.height = videoRef.current.videoHeight;
    canvas.getContext('2d')?.drawImage(videoRef.current, 0, 0);
    const dataUrl = canvas.toDataURL('image/jpeg');
    updatePhotoState(activeCamera, dataUrl);
    if (activeCamera !== 'kitchen' && activeCamera !== 'meal') {
      setActiveCamera(null);
    }
  };

  const updatePhotoState = (type: string, dataUrl: string) => {
    setCapturedPhotos(prev => {
      const current = prev[type] || [];
      if (type === 'exterior' || type === 'selfie') return { ...prev, [type]: [dataUrl] };
      return { ...prev, [type]: [...current, dataUrl] };
    });
  };

  const removePhoto = (type: string, index: number) => {
    setCapturedPhotos(prev => ({
      ...prev,
      [type]: prev[type].filter((_, i) => i !== index)
    }));
  };

  const getLocation = () => {
    setGeoStatus('loading');
    navigator.geolocation.getCurrentPosition(
      (pos) => {
        updateProof({ location: { lat: pos.coords.latitude, lng: pos.coords.longitude } });
        setGeoStatus('success');
      },
      () => setGeoStatus('error')
    );
  };

  const handleNext = () => {
    updateProof({
      kitchenPhotos: capturedPhotos.kitchen,
      mealPhotos: capturedPhotos.meal,
      exteriorPhoto: capturedPhotos.exterior[0] || null,
      selfie: capturedPhotos.selfie[0] || null
    });
    setStep(4);
  };

  return (
    <div className="p-8 space-y-8 relative overflow-hidden">
      <div className="space-y-1">
        <h2 className="text-xl font-bold text-dark">Proof of Life Verification</h2>
        <p className="text-sm text-neutral">Live evidence of your active kitchen operation.</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div className="space-y-6">
          {/* Photo Sections */}
          {[
            { id: 'kitchen', label: 'Kitchen Environment', count: 2, desc: 'Pantry, stove, and prep area' },
            { id: 'meal', label: 'Prepared Meals', count: 3, desc: 'Freshly cooked meal samples' },
            { id: 'exterior', label: 'Exterior Shop/House', count: 1, desc: 'View from the street' }
          ].map(section => (
            <div key={section.id} className="space-y-3">
              <div className="flex items-center justify-between">
                <label className="text-xs font-bold text-neutral uppercase tracking-wider">{section.label}</label>
                <span className="text-[10px] font-bold text-neutral/50">{capturedPhotos[section.id]?.length || 0} / {section.count}</span>
              </div>
              <div className="flex flex-wrap gap-2">
                {capturedPhotos[section.id]?.map((img, i) => (
                  <div key={i} className="relative w-16 h-16 rounded-lg overflow-hidden group border border-slate-100">
                    <img src={img} className="w-full h-full object-cover" />
                    <button 
                      onClick={() => removePhoto(section.id, i)}
                      className="absolute inset-0 bg-black/40 opacity-0 group-hover:opacity-100 flex items-center justify-center transition-opacity"
                    >
                      <Icon icon="lucide:x" className="w-4 h-4 text-white" />
                    </button>
                  </div>
                ))}
                {capturedPhotos[section.id]?.length < section.count && (
                  <button 
                    onClick={() => setActiveCamera(section.id as any)}
                    className="w-16 h-16 rounded-lg border-2 border-dashed border-slate-200 flex items-center justify-center hover:border-primary hover:bg-primary/5 transition-all text-slate-400 hover:text-primary"
                  >
                    <Icon icon="lucide:camera" className="w-6 h-6" />
                  </button>
                )}
              </div>
            </div>
          ))}

          {/* Smart Selfie Section */}
          <div className="space-y-3 p-4 bg-primary/5 rounded-2xl border border-primary/10">
            <div className="flex items-center justify-between">
              <label className="text-xs font-bold text-primary uppercase tracking-wider flex items-center gap-2">
                 <Icon icon="lucide:shield-check" className="w-3.5 h-3.5" /> Smile ID Liveness
              </label>
            </div>
            <div className="flex items-center gap-4">
              {capturedPhotos.selfie[0] ? (
                <div className="w-20 h-20 rounded-full border-2 border-primary overflow-hidden relative">
                  <img src={capturedPhotos.selfie[0]} className="w-full h-full object-cover" />
                  <button onClick={() => removePhoto('selfie', 0)} className="absolute inset-0 bg-black/40 flex items-center justify-center opacity-0 hover:opacity-100">
                    <Icon icon="lucide:refresh-ccw" className="w-4 h-4 text-white" />
                  </button>
                </div>
              ) : (
                <button 
                  onClick={() => setActiveCamera('selfie')}
                  className="w-20 h-20 rounded-full border-2 border-dashed border-primary/30 flex items-center justify-center hover:bg-primary/10 transition-all text-primary"
                >
                  <Icon icon="lucide:camera" className="w-8 h-8" />
                </button>
              )}
              <div className="flex-1 space-y-1">
                <p className="text-xs font-bold text-dark">Liveness & Identity Check</p>
                <p className="text-[10px] text-neutral">Powered by Smile ID to prevent ghost vendors.</p>
              </div>
            </div>
          </div>
        </div>

        {/* Location Section */}
        <div className="space-y-6">
          <div className="space-y-3">
             <label className="text-xs font-bold text-neutral uppercase tracking-wider block">Live Geolocation</label>
             <div className="h-48 bg-slate-50 rounded-2xl flex flex-col items-center justify-center border border-slate-200 relative overflow-hidden">
                {data.proof.location ? (
                  <div className="absolute inset-0 flex flex-col items-center justify-center bg-white">
                    <Icon icon="lucide:map-pin" className="w-8 h-8 text-primary animate-bounce mb-2" />
                    <p className="text-[10px] font-bold text-dark">LOCATION PINNED</p>
                    <p className="text-[10px] text-neutral">{data.proof.location.lat.toFixed(4)}, {data.proof.location.lng.toFixed(4)}</p>
                    <button onClick={getLocation} className="mt-2 text-[10px] font-bold text-primary hover:underline">Repin</button>
                  </div>
                ) : (
                  <>
                    <Icon icon="lucide:map-pin" className="w-8 h-8 text-slate-300 mb-2" />
                    <button 
                      onClick={getLocation}
                      disabled={geoStatus === 'loading'}
                      className="bg-white px-4 py-2 rounded-xl text-xs font-bold shadow-sm hover:shadow-md transition-all flex items-center gap-2"
                    >
                      {geoStatus === 'loading' ? 'Fetching...' : 'Verify Location'}
                      <Icon icon="lucide:upload" className="w-3 h-3" />
                    </button>
                  </>
                )}
             </div>
             <p className="text-[10px] text-neutral/60 italic">This must match your business address for verification.</p>
          </div>
        </div>
      </div>

      <div className="flex gap-4">
        <button
          type="button"
          onClick={() => setStep(2)}
          className="flex-1 bg-slate-50 text-neutral rounded-xl py-4 font-bold border border-slate-100 hover:bg-slate-100 transition-all flex items-center justify-center gap-2"
        >
          <Icon icon="lucide:arrow-left" className="w-5 h-5" /> Back
        </button>
        <button
          onClick={handleNext}
          disabled={!capturedPhotos.selfie[0] || !data.proof.location || capturedPhotos.kitchen.length < 2 || capturedPhotos.meal.length < 3}
          className="flex-[2] bg-dark text-white rounded-xl py-4 font-bold hover:opacity-90 transition-all flex items-center justify-center gap-2 shadow-lg disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Review & Submit
          <Icon icon="lucide:arrow-right" className="w-5 h-5" />
        </button>
      </div>

      {/* Camera Modal */}
      <AnimatePresence>
        {activeCamera && (
          <motion.div 
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            className="fixed inset-0 z-[60] bg-black flex items-center justify-center p-4"
          >
            <div className="relative w-full max-w-lg aspect-[3/4] bg-dark rounded-3xl overflow-hidden shadow-2xl">
              <video ref={videoRef} autoPlay playsInline className="w-full h-full object-cover" />
              
              <div className="absolute inset-0 pointer-events-none">
                 {activeCamera === 'selfie' && (
                   <div className="absolute inset-0 flex items-center justify-center">
                      <div className="w-64 h-80 border-2 border-primary/50 rounded-[100px] shadow-[0_0_0_1000px_rgba(0,0,0,0.5)]" />
                   </div>
                 )}
              </div>

              <div className="absolute top-6 left-0 right-0 flex justify-center text-white/80 text-xs font-medium uppercase tracking-[0.2em]">
                {activeCamera === 'selfie' ? 'Position face in frame' : `Capturing ${activeCamera}`}
              </div>

              <button 
                onClick={() => setActiveCamera(null)}
                className="absolute top-6 right-6 w-10 h-10 rounded-full bg-white/20 backdrop-blur flex items-center justify-center text-white"
              >
                <Icon icon="lucide:x" className="w-6 h-6" />
              </button>

              <div className="absolute bottom-8 left-0 right-0 flex flex-col items-center gap-4">
                {countdown !== null && (
                  <motion.div 
                    initial={{ scale: 0.5 }}
                    animate={{ scale: 1.5 }}
                    className="text-6xl font-bold text-primary"
                  >
                    {countdown}
                  </motion.div>
                )}
                <button 
                  onClick={capturePhoto}
                  className="w-20 h-20 rounded-full border-4 border-white flex items-center justify-center group"
                >
                  <div className="w-16 h-16 rounded-full bg-white group-active:scale-90 transition-transform" />
                </button>
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
}
