package db

import (
	"context"
	"era/booru/ent"
	"era/booru/ent/tag"
	"fmt"
)

func FindOrCreateTag(ctx context.Context, db *ent.Client, name string) (*ent.Tag, error) {
	tg, err := db.Tag.Query().Where(tag.NameEQ(name)).Only(ctx)
	if ent.IsNotFound(err) {
		tg, err = db.Tag.Create().SetName(name).SetType(tag.TypeUserTag).Save(ctx)
	}
	return tg, err
}

// Alternative: if you want to handle multiple tags at once
func FindOrCreateTags(ctx context.Context, db *ent.Client, tagNames []string) ([]int, error) {
	tagIDs := make([]int, 0, len(tagNames))
	for _, name := range tagNames {
		tg, err := FindOrCreateTag(ctx, db, name)
		if err != nil {
			return nil, fmt.Errorf("tag lookup/create %s: %w", name, err)
		}
		tagIDs = append(tagIDs, tg.ID)
	}
	return tagIDs, nil
}

// SetMediaTags replaces all tags on the given media item with the provided list.
// The tags slice should already be normalized (trimmed and deduplicated).
func SetMediaTags(ctx context.Context, db *ent.Client, mediaID string, tags []string) error {
	tagIDs, err := FindOrCreateTags(ctx, db, tags)
	if err != nil {
		return err
	}
	_, err = db.Media.UpdateOneID(mediaID).ClearTags().AddTagIDs(tagIDs...).Save(ctx)
	return err
}
