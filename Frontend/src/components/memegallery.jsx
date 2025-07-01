import React, { useState } from "react";

const MemeGallery = () => {
  const dummyMemes = [
    {
      id: 1,
      title: "Grumpy Cat",
      image_url: "https://picsum.photos/200",
      tags: ["cat", "funny", "grumpy"],
    },
    {
      id: 2,
      title: "Distracted Boyfriend",
      image_url: "https://picsum.photos/200",
      tags: ["meme", "funny", "relationship"],
    },
    {
      id: 3,
      title: "Success Kid",
      image_url: "https://picsum.photos/200",
      tags: ["kid", "success", "funny"],
    },
    {
      id: 4,
      title: "Doge",
      image_url: "https://picsum.photos/200",
      tags: ["dog", "doge", "cute"],
    },
    {
      id: 5,
      title: "Drake Hotline Bling",
      image_url: "https://picsum.photos/200",
      tags: ["drake", "funny", "reaction"],
    },
    {
      id: 6,
      title: "Bad Luck Brian",
      image_url: "https://picsum.photos/200",
      tags: ["funny", "meme", "unlucky"],
    },
  ];

  return (
    <div
      style={{
        display: "grid",
        gridTemplateColumns: "repeat(3, 1fr)",
        gap: "16px",
        padding: "16px",
      }}
    >
      {dummyMemes.length === 0 ? (
        <p style={{ gridColumn: "span 3", textAlign: "center" }}>
          No memes yet. Add some!
        </p>
      ) : (
        dummyMemes.map((meme) => (
          <div
            key={meme.id}
            style={{
              border: "1px solid #ccc",
              borderRadius: "4px",
              overflow: "hidden",
              height: "280px",
              display: "flex",
              flexDirection: "column",
            }}
          >
            <img
              src={meme.image_url}
              alt={meme.title}
              style={{
                width: "100%",
                height: "100px",
                objectFit: "cover",
                flexShrink: 0,
              }}
              onError={(e) => {
                e.target.src =
                  "https://via.placeholder.com/300?text=Image+Not+Found";
              }}
            />
            <div
              style={{
                padding: "8px",
                flexGrow: 1,
                display: "flex",
                flexDirection: "column",
                justifyContent: "space-between",
              }}
            >
              <div>
                <h3 style={{ fontSize: "14px", marginBottom: "4px" }}>
                  {meme.title}
                </h3>

                <input
                  type="number"
                  placeholder="Enter bid amount"
                  style={{
                    width: "100%",
                    padding: "4px",
                    fontSize: "12px",
                    marginBottom: "8px",
                  }}
                />
              </div>
              <button
                style={{
                  padding: "6px",
                  fontSize: "12px",
                  cursor: "pointer",
                }}
              >
                Submit Bid
              </button>
            </div>
          </div>
        ))
      )}
    </div>
  );
};

export default MemeGallery;
